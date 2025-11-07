# VM Errors, Snapshots, and State Rollback: Complete Analysis

**Date**: September 25, 2024  
**Context**: Deep dive into Ethereum VM error handling, state rollback mechanisms, and execution flow

## Executive Summary

This analysis provides a rigorous proof that VM errors in Ethereum's execution model are properly handled through a sophisticated snapshot-based rollback system. The key insight is the fundamental distinction between **consensus errors** (which prevent transaction inclusion) and **VM execution errors** (which allow inclusion but mark execution as failed).

## Key Questions Addressed

1. **When do VM errors occur during contract calls?**
2. **What constitutes a failed call that stops execution?**
3. **How does state rollback work when VM execution fails?**
4. **Do VM errors actually stop execution, and can execution continue after rollback?**
5. **How are VM errors properly labeled and distinguished from consensus errors?**

## Rigorous Proof: VM Error Call Chain

### 1. Top-Level Transaction Processing

```go
// From core/state_processor.go:154
result, err := ApplyMessage(evm, msg, gp)
if err != nil {
    return nil, err  // ← CONSENSUS ERROR (transaction rejected)
}
```

### 2. State Transition Level

```go
// From core/state_transition.go:189-190
func ApplyMessage(evm *vm.EVM, msg *Message, gp *GasPool) (*ExecutionResult, error) {
    return NewStateTransition(evm, msg, gp).TransitionDb()
}
```

### 3. VM Execution with VM Error Capture

```go
// From core/state_transition.go:445-455
var (
    ret   []byte
    vmerr error // vm errors do not effect consensus and are therefore not assigned to err
)
if contractCreation {
    ret, _, st.gasRemaining, vmerr = st.evm.Create(sender, msg.Data, st.gasRemaining, value)
} else {
    // Increment the nonce for the next transaction
    st.state.SetNonce(msg.From, st.state.GetNonce(sender.Address())+1)
    ret, st.gasRemaining, vmerr = st.evm.Call(sender, st.to(), msg.Data, st.gasRemaining, value)
}
```

**CRITICAL PROOF**: The error from `st.evm.Call()` is assigned to `vmerr` (VM error), NOT to the consensus error.

### 4. ExecutionResult Construction

```go
// From core/state_transition.go:489-494
return &ExecutionResult{
    UsedGas:     st.gasUsed(),
    RefundedGas: gasRefund,
    Err:         vmerr,  // ← VM ERROR STORED HERE
    ReturnData:  ret,
}, nil  // ← CONSENSUS ERROR IS ALWAYS NIL
```

### 5. Receipt Status Determination

```go
// From core/state_processor.go:176-180
if result.Failed() {  // ← result.Failed() checks result.Err != nil
    receipt.Status = types.ReceiptStatusFailed
} else {
    receipt.Status = types.ReceiptStatusSuccessful
}
```

## State Rollback Mechanism: Snapshot-Based System

### How Snapshots Work

The EVM uses a **snapshot-based rollback system** implemented through the `StateDB` interface:

```go
// From core/vm/interface.go:91-92
RevertToSnapshot(int)
Snapshot() int
```

### Snapshot Lifecycle

1. **Before Each Call - Take a Snapshot**:
```go
// From core/vm/evm.go:188 (Call method)
snapshot := evm.StateDB.Snapshot()

// From core/vm/evm.go:283 (DelegateCall method)  
var snapshot = evm.StateDB.Snapshot()

// From core/vm/evm.go:516 (Create method)
snapshot := evm.StateDB.Snapshot()
```

2. **During Execution - State Changes Accumulate**:
   - Balance transfers
   - Storage modifications  
   - Code deployment
   - Account creation
   - Nonce increments

3. **If Execution Fails - Revert to Snapshot**:
```go
// From core/vm/evm.go:307-316 (Call method)
if err != nil {
    evm.StateDB.RevertToSnapshot(snapshot)  // ← ROLLBACK HAPPENS HERE
    if err != ErrExecutionReverted {
        if evm.Config.Tracer != nil && evm.Config.Tracer.OnGasChange != nil {
            evm.Config.Tracer.OnGasChange(gas, 0, tracing.GasChangeCallFailedExecution)
        }
        gas = 0
    }
}
```

### What Gets Rolled Back

When `RevertToSnapshot(snapshot)` is called, **ALL** state changes made during that call are undone:

- ✅ **Balance transfers** - reverted
- ✅ **Storage modifications** - reverted  
- ✅ **Code deployment** - reverted
- ✅ **Account creation** - reverted
- ✅ **Nonce increments** - reverted
- ✅ **Access list changes** - reverted
- ✅ **Refund counter changes** - reverted

### What Does NOT Get Rolled Back

However, some things are **NOT** rolled back because they happen at the **state transition level**:

- ❌ **Gas consumption** - gas is still charged for the failed execution
- ❌ **Transaction nonce increment** - happens in `state_transition.go:453`
- ❌ **Gas refunds** - processed regardless of VM error
- ❌ **Fee payments to miner** - still paid

## VM Error Types and Execution Termination

### Pre-execution VM Errors (Immediate Return)

These errors occur **before** any state changes and cause **immediate return**:

```go
// From core/vm/evm.go:181-182 (Call method)
if evm.depth > int(params.CallCreateDepth) {
    return nil, gas, ErrDepth  // ← EXECUTION STOPS IMMEDIATELY
}

// From core/vm/evm.go:185-186 (Call method)  
if !value.IsZero() && !evm.Context.CanTransfer(evm.StateDB, caller.Address(), value) {
    return nil, gas, ErrInsufficientBalance  // ← EXECUTION STOPS IMMEDIATELY
}
```

### Post-snapshot VM Errors (Revert + Return)

These errors occur **after** a snapshot is taken and cause **revert + return**:

```go
// From core/vm/evm.go:195-198 (Call method)
if gas < wgas {
    evm.StateDB.RevertToSnapshot(snapshot)  // ← REVERT STATE
    return nil, 0, ErrOutOfGas              // ← EXECUTION STOPS
}
```

### Interpreter VM Errors (Immediate Return from Interpreter)

The interpreter's main loop **immediately returns** on any VM error:

```go
// From core/vm/interpreter.go:253-254
if contract.Gas < cost {
    return nil, ErrOutOfGas  // ← INTERPRETER STOPS IMMEDIATELY
}

// From core/vm/interpreter.go:247-248
if sLen := stack.len(); sLen < operation.minStack {
    return nil, &ErrStackUnderflow{stackLen: sLen, required: operation.minStack}  // ← INTERPRETER STOPS
}
```

## Nested Call Error Handling

### CALL Opcode Behavior

When a contract makes a nested call, the CALL opcode handles VM errors:

```go
// From core/vm/instructions.go:756-771
ret, returnGas, err := interpreter.evm.Call(scope.Contract, toAddr, args, gas, &value)

if err != nil {
    temp.Clear()  // ← Push 0 to stack (call failed)
} else {
    temp.SetOne() // ← Push 1 to stack (call succeeded)
}
stack.push(&temp)
if err == nil || err == ErrExecutionReverted {
    scope.Memory.Set(retOffset.Uint64(), retSize.Uint64(), ret)
}
```

**KEY INSIGHT**: The CALL opcode **does NOT stop execution** when it receives a VM error. It:
- Pushes 0 to the stack (indicating failure)
- Stores return data if available
- **Continues execution** of the calling contract

## Critical Distinctions

### Consensus Errors vs VM Errors

#### **Consensus Errors** (Stop Transaction Inclusion):
```go
// From core/state_transition.go:394-396
if err := st.preCheck(); err != nil {
    return nil, err  // ← CONSENSUS ERROR - transaction rejected
}
```

**Examples:**
- Invalid nonce (`ErrNonceTooLow/High`)
- Insufficient balance for gas + value (`ErrInsufficientFunds`)
- Gas limit exceeded (`ErrGasLimitReached`)
- Intrinsic gas too low (`ErrIntrinsicGas`)

#### **VM Errors** (Allow Inclusion, Mark as Failed):
```go
// From core/state_transition.go:489-494
return &ExecutionResult{
    Err: vmerr,  // ← VM ERROR - transaction included but marked failed
}, nil  // ← CONSENSUS ERROR IS NIL
```

**Examples:**
- `ErrOutOfGas` - Insufficient gas during opcode execution
- `ErrExecutionReverted` - REVERT opcode executed
- `ErrDepth` - Call stack exceeded 1024 levels
- `ErrInsufficientBalance` - Contract balance < transfer value

## Practical Transaction Outcomes

| Scenario | Consensus Error | VM Error | Block Inclusion | Receipt Status | Gas Charged |
|----------|-----------------|----------|-----------------|----------------|-------------|
| Invalid nonce | ✓ | - | ❌ Rejected | No receipt | ❌ No gas |
| Out of gas | - | ✓ | ✅ Included | Failed (0) | ✅ Up to failure point |
| Successful execution | - | - | ✅ Included | Successful (1) | ✅ Total used |
| Revert execution | - | ✓ | ✅ Included | Failed (0) | ✅ Total used |

## Key Takeaways

### 1. **VM Errors DO Stop Execution**
- **Interpreter has no error recovery** - it returns immediately on any error
- **EVM has no error recovery** - it reverts and returns immediately on any error  
- **No continuation mechanism exists** - there's no way to "resume" after a VM error

### 2. **State Rollback is Guaranteed**
- **Snapshots are call-scoped** - each call gets its own snapshot
- **All state changes are reverted** on VM errors
- **Gas is still charged** - computational work is always paid for

### 3. **Nested Calls Don't Stop Parent Execution**
- **CALL opcode continues execution** after receiving VM errors
- **Parent contracts can handle** child call failures
- **Only top-level VM errors** affect transaction receipt status

### 4. **Architectural Guarantees**
- **Atomicity**: Each call is atomic (all-or-nothing)
- **Isolation**: Calls don't interfere with each other's state
- **Consistency**: Failed calls don't leave partial state changes
- **Durability**: Successful state changes are permanent

### 5. **Resource Accountability**
- **All computational work is paid for** - preventing DoS attacks
- **Failed transactions still consume block space and gas**
- **Deterministic gas costs** ensure predictable network behavior

## Conclusion

The Ethereum VM error architecture demonstrates sophisticated design principles that separate consensus validation from execution outcomes. **VM errors like `ErrOutOfGas` are execution errors, not consensus errors** - they allow transaction inclusion while marking execution as failed, ensuring computational work is always paid for while maintaining blockchain consistency and miner incentives.

This design enables:
1. **Predictable resource costs** for users
2. **State consistency** across the network  
3. **Miner compensation** for computational work
4. **Consensus safety** through deterministic execution

The distinction between "transaction success" (inclusion in block) and "execution success" (receipt status) is fundamental to understanding Ethereum's execution model.

## Validation Against Original Notes

The analysis in `geth-vm-errors.md` was **remarkably accurate** and well-aligned with the actual codebase:

- ✅ **VM Error Definitions** - All 13 core VM errors match exactly
- ✅ **Consensus vs VM Error Architecture** - Fundamental distinction correctly captured
- ✅ **ExecutionResult Structure** - Matches exactly with implementation
- ✅ **Key Methods** - All methods exist and work as described
- ✅ **Critical Architecture Insight** - Key pattern correctly identified
- ✅ **Receipt Status Logic** - Matches exactly with implementation
- ✅ **State Transition Return** - Final return statement confirms key insight

The original notes serve as an excellent reference for understanding Ethereum's VM error architecture and demonstrate a sophisticated understanding of the underlying design principles.
