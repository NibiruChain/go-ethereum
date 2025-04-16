// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package native

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/params"
)

func init() {
	tracers.DefaultDirectory.Register("noopTracer", NewNoopTracer, false)
}

// NoOpTracer is a go implementation of the Tracer interface which
// performs no action. It's mostly useful for testing purposes.
type NoOpTracer struct{}

// NewNoopTracer returns a new noop tracer.
func NewNoopTracer(ctx *tracers.Context, cfg json.RawMessage, chainConfig *params.ChainConfig) (*tracers.Tracer, error) {
	t := &NoOpTracer{}
	return &tracers.Tracer{
		Hooks: &tracing.Hooks{
			OnTxStart:       t.OnTxStart,
			OnTxEnd:         t.OnTxEnd,
			OnEnter:         t.OnEnter,
			OnExit:          t.OnExit,
			OnOpcode:        t.OnOpcode,
			OnFault:         t.OnFault,
			OnGasChange:     t.OnGasChange,
			OnBalanceChange: t.OnBalanceChange,
			OnNonceChange:   t.OnNonceChange,
			OnCodeChange:    t.OnCodeChange,
			OnStorageChange: t.OnStorageChange,
			OnLog:           t.OnLog,
		},
		GetResult: t.GetResult,
		Stop:      t.Stop,
	}, nil
}

func (t *NoOpTracer) OnOpcode(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, rData []byte, depth int, err error) {
}

func (t *NoOpTracer) OnFault(pc uint64, op byte, gas, cost uint64, _ tracing.OpContext, depth int, err error) {
}

func (t *NoOpTracer) OnGasChange(old, new uint64, reason tracing.GasChangeReason) {}

func (t *NoOpTracer) OnEnter(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
}

func (t *NoOpTracer) OnExit(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
}

func (*NoOpTracer) OnTxStart(env *tracing.VMContext, tx *types.Transaction, from common.Address) {
}

func (*NoOpTracer) OnTxEnd(receipt *types.Receipt, err error) {}

func (*NoOpTracer) OnBalanceChange(a common.Address, prev, new *big.Int, reason tracing.BalanceChangeReason) {
}

func (*NoOpTracer) OnNonceChange(a common.Address, prev, new uint64) {}

func (*NoOpTracer) OnCodeChange(a common.Address, prevCodeHash common.Hash, prev []byte, codeHash common.Hash, code []byte) {
}

func (*NoOpTracer) OnStorageChange(a common.Address, k, prev, new common.Hash) {}

func (*NoOpTracer) OnLog(log *types.Log) {}

// GetResult returns an empty json object.
func (t *NoOpTracer) GetResult() (json.RawMessage, error) {
	return json.RawMessage(`{}`), nil
}

// Stop terminates execution of the tracer at the first opportune moment.
func (t *NoOpTracer) Stop(err error) {
}
