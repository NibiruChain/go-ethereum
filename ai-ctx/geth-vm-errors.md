# Comprehensive Analysis: All Possible VM Errors in Ethereum State Transitions

Based on rigorous analysis of the Nibiru-Geth codebase, I've identified a clear architectural distinction between **consensus errors** (which prevent transaction inclusion) and **VM execution errors** (which allow inclusion but mark execution as failed). Here's a comprehensive characterization of all VM errors.

## Executive Summary

The Ethereum execution model maintains a sharp separation between:
- **Consensus Layer**: Validates transaction validity and block constraints
- **Execution Layer**: Handles VM execution and computational errors

**Key Insight**: Transactions can be "processed successfully" (included in blocks) while being marked as "failed" (receipt status = 0) when VM execution errors occur.

## 1. VM Error Definitions and Classification

### Core VM Errors (`core/vm/errors.go:27-45`)

| Error | Code | Definition | Trigger Point |
|-------|------|------------|---------------|
| `ErrOutOfGas` | 1 | "out of gas" | Insufficient gas during opcode execution |
| `ErrCodeStoreOutOfGas` | 2 | "contract creation code storage out of gas" | Contract code storage exceeds gas |
| `ErrDepth` | 3 | "max call depth exceeded" | Call stack > 1024 levels |
| `ErrInsufficientBalance` | 4 | "insufficient balance for transfer" | Contract balance < transfer value |
| `ErrContractAddressCollision` | 5 | "contract address collision" | CREATE at existing address |
| `ErrExecutionReverted` | 6 | "execution reverted" | REVERT opcode executed |
| `ErrMaxCodeSizeExceeded` | 7 | "max code size exceeded" | Contract code > 24KB |
| `ErrInvalidJump` | 8 | "invalid jump destination" | JUMP to invalid destination |
| `ErrWriteProtection` | 9 | "write protection" | State modification in static call |
| `ErrReturnDataOutOfBounds` | 10 | "return data out of bounds" | RETURNDATASIZE/COPY bounds violation |
| `ErrGasUintOverflow` | 11 | "gas uint64 overflow" | Gas calculation overflow |
| `ErrInvalidCode` | 12 | "invalid code: must not begin with 0xef" | Invalid EOF prefix |
| `ErrNonceUintOverflow` | 13 | "nonce uint64 overflow" | Account nonce overflow |

### Dynamic VM Errors (`core/vm/errors.go:49-83`)

| Error | Code | Description |
|-------|------|-------------|
| `ErrStackUnderflow` | 14 | Stack has fewer items than required |
| `ErrStackOverflow` | 15 | Stack exceeds 1024 item limit |
| `ErrInvalidOpCode` | 16 | Undefined opcode encountered |

### Internal Token
- `errStopToken`: Internal interpreter termination signal (never returned to callers)

## 2. VM Error Trigger Locations in State Transitions

### Primary Trigger Points

#### A. Interpreter Main Loop (`core/vm/interpreter.go:253-254`)
```go
if contract.Gas < cost {
    return nil, ErrOutOfGas
}
```
**Analysis**: Before each opcode execution, checks if sufficient gas remains for constant gas cost.

#### B. Dynamic Gas Operations (`core/vm/interpreter.go:286-287`)
```go
if contract.Gas < dynamicCost {
    return nil, ErrOutOfGas
}
```
**Analysis**: After computing dynamic gas costs (memory expansion, storage operations).

#### C. EVM Call/Create Operations (`core/vm/evm.go:182,274,339,392,470`)
```go
if depth > int(params.CallCreateDepth) {
    return nil, gas, ErrDepth
}
```
**Analysis**: Call depth validation across all EVM operations.

#### D. Balance Validation (`core/vm/evm.go:186,281,473`)
```go
if !evm.Context.CanTransfer(evm.StateDB, caller.Address(), value) {
    return nil, gas, ErrInsufficientBalance
}
```
**Analysis**: Validates sufficient balance for transfers.

#### E. Contract Creation Validation (`core/vm/evm.go:568-583`)
```go
if evm.chainRules.IsShanghai && len(ret) > params.MaxCodeSize {
    return ret, ErrMaxCodeSizeExceeded
}
```
**Analysis**: Validates contract code size limits and EOF compliance.

## 3. Consensus vs Execution Error Architecture

### Consensus Errors (`core/error.go:42-115`)
**Effect**: Prevent transaction inclusion in blocks

| Error | Description | Returns |
|-------|-------------|---------|
| `ErrNonceTooLow/High` | Invalid nonce | Consensus Error |
| `ErrInsufficientFunds` | Insufficient balance for gas+value | Consensus Error |
| `ErrGasLimitReached` | Block gas limit exceeded | Consensus Error |
| `ErrIntrinsicGas` | Gas below minimum required | Consensus Error |
| `ErrMaxInitCodeSizeExceeded` | Init code too large | Consensus Error |

### VM Execution Errors (`core/vm/errors.go`)
**Effect**: Allow transaction inclusion, mark execution as failed

**Critical Architecture**: `core/state_transition.go:445-447`
```go
var (
    ret   []byte
    vmerr error // vm errors do not effect consensus and are therefore not assigned to err
)
```

## 4. Error Propagation Through ExecutionResult

### ExecutionResult Structure (`core/state_transition.go:25-33`)
```go
type ExecutionResult struct {
    UsedGas     uint64 // Total used gas, not including the refunded gas
    RefundedGas uint64 // Total gas refunded after execution
    Err         error  // Any error encountered during the execution(listed in core/vm/errors.go)
    ReturnData  []byte // Returned data from evm(function result or data supplied with revert opcode)
}
```

### Key Methods
- `Failed() bool`: Returns `result.Err != nil`
- `Unwrap() error`: Returns the VM error
- `Return() []byte`: Returns data if no error
- `Revert() []byte`: Returns revert data if `ErrExecutionReverted`

### Receipt Status Determination (`core/state_processor.go:176-180`)
```go
if result.Failed() {
    receipt.Status = types.ReceiptStatusFailed
} else {
    receipt.Status = types.ReceiptStatusSuccessful
}
```

## 5. State Transition Flow with VM Errors

### Transaction Processing Flow
1. **Pre-execution Checks**: Validate nonce, balance, gas limits
2. **Gas Purchase**: Deduct gas costs from sender balance
3. **VM Execution**: Execute transaction, capture VM errors in `vmerr`
4. **Error Handling**: VM errors stored but execution continues
5. **Gas Refunds**: Processed regardless of VM error status
6. **Fee Payment**: Miners receive fees for consumed gas
7. **Result Construction**: Return `ExecutionResult{Err: vmerr}, nil`

### Critical Behavior (`core/state_transition.go:489-494`)
```go
return &ExecutionResult{
    UsedGas:     st.gasUsed(),
    RefundedGas: gasRefund,
    Err:         vmerr,  // VM error (may be ErrOutOfGas)
    ReturnData:  ret,
}, nil  // CONSENSUS ERROR IS NIL
```

**Key Insight**: State transition always returns `nil` consensus error when VM execution completes, regardless of VM errors.

## 6. Gas and Fee Handling with VM Errors

### Gas Consumption Despite VM Errors
- Gas charged for all computational work before failure
- Refunds processed regardless of execution outcome  
- Miners receive fees for consumed gas
- Remaining gas returned to sender and block pool

### Gas Calculation (`core/state_transition.go:527-529`)
```go
func (st *StateTransition) gasUsed() uint64 {
    return st.initialGas - st.gasRemaining
}
```

## 7. Practical Transaction Outcomes

| Scenario | Consensus Error | VM Error | Block Inclusion | Receipt Status | Gas Charged |
|----------|-----------------|----------|-----------------|----------------|-------------|
| Invalid nonce | ✓ | - | ❌ Rejected | No receipt | ❌ No gas |
| Out of gas | - | ✓ | ✅ Included | Failed (0) | ✅ Up to failure point |
| Successful execution | - | - | ✅ Included | Successful (1) | ✅ Total used |
| Revert execution | - | ✓ | ✅ Included | Failed (0) | ✅ Total used |

## 8. Security and Consistency Implications

### Resource Accountability
- All computational work is paid for, preventing DoS attacks
- Failed transactions still consume block space and gas
- Deterministic gas costs ensure predictable network behavior

### State Consistency  
- VM errors preserve state changes made before failure
- Nonce incrementation occurs regardless of execution outcome
- Gas refunds maintain economic incentives

### Consensus Safety
- VM errors are deterministic and reproducible across nodes
- Consensus errors prevent non-deterministic transaction inclusion
- Clear separation prevents execution issues from affecting consensus

## Conclusion

Ethereum's VM error architecture demonstrates sophisticated design principles that separate consensus validation from execution outcomes. **VM errors like `ErrOutOfGas` are execution errors, not consensus errors** - they allow transaction inclusion while marking execution as failed, ensuring computational work is always paid for while maintaining blockchain consistency and miner incentives.

This design enables:
1. **Predictable resource costs** for users
2. **State consistency** across the network  
3. **Miner compensation** for computational work
4. **Consensus safety** through deterministic execution

The distinction between "transaction success" (inclusion in block) and "execution success" (receipt status) is fundamental to understanding Ethereum's execution model.