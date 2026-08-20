package vm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

type trueCallerCapture struct {
	address    common.Address
	trueCaller common.Address
}

func (p *trueCallerCapture) RequiredGas([]byte) uint64 { return 0 }

func (p *trueCallerCapture) Address() common.Address { return p.address }

func (p *trueCallerCapture) Run(
	_ *EVM,
	trueCaller common.Address,
	_ *Contract,
	_ bool,
	_ bool,
) ([]byte, error) {
	p.trueCaller = trueCaller
	return nil, nil
}

func TestTrueCallerPropagation(t *testing.T) {
	user := common.HexToAddress("0x1000000000000000000000000000000000000001")
	proxy := common.HexToAddress("0x2000000000000000000000000000000000000002")
	implementation := common.HexToAddress("0x3000000000000000000000000000000000000003")
	pool := common.HexToAddress("0x4000000000000000000000000000000000000004")
	attacker := common.HexToAddress("0x5000000000000000000000000000000000000005")

	root := NewContract(AccountRef(user), AccountRef(proxy), uint256.NewInt(0), 1_000_000)
	root.TrueCaller = user
	if root.TrueCaller != user {
		t.Fatalf("root true caller: got %s want %s", root.TrueCaller, user)
	}

	proxyDelegate := NewContract(root, AccountRef(proxy), nil, 1_000_000).AsDelegate()
	proxyDelegate.SetCallCode(&implementation, common.Hash{}, nil)
	if proxyDelegate.TrueCaller != user {
		t.Fatalf("delegatecall did not preserve true caller: got %s want %s", proxyDelegate.TrueCaller, user)
	}

	poolFrame := NewContract(root, AccountRef(pool), uint256.NewInt(0), 1_000_000)
	if poolFrame.TrueCaller != pool {
		t.Fatalf("CALL did not reset true caller to pool: got %s want %s", poolFrame.TrueCaller, pool)
	}

	callbackFrame := NewContract(poolFrame, AccountRef(attacker), uint256.NewInt(0), 1_000_000)
	if callbackFrame.TrueCaller != attacker {
		t.Fatalf("callback CALL did not reset true caller to attacker: got %s want %s", callbackFrame.TrueCaller, attacker)
	}
}

func TestDelegateCallPassesTrueCallerToPrecompile(t *testing.T) {
	pool := common.HexToAddress("0x4000000000000000000000000000000000000004")
	attacker := common.HexToAddress("0x5000000000000000000000000000000000000005")
	precompileAddr := common.HexToAddress("0x0000000000000000000000000000000000000800")

	statedb, err := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())
	if err != nil {
		t.Fatal(err)
	}
	evm := NewEVM(BlockContext{}, TxContext{}, statedb, params.TestChainConfig, Config{})
	capture := &trueCallerCapture{address: precompileAddr}
	evm.SetPrecompiles(PrecompiledContracts{precompileAddr: capture})

	callbackFrame := NewContract(AccountRef(pool), AccountRef(attacker), uint256.NewInt(0), 1_000_000)
	if _, _, err := evm.DelegateCall(callbackFrame, precompileAddr, nil, 1_000_000); err != nil {
		t.Fatal(err)
	}
	if capture.trueCaller != attacker {
		t.Fatalf("precompile true caller: got %s want callback contract %s", capture.trueCaller, attacker)
	}
}
