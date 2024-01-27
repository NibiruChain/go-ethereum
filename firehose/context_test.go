package firehose

import (
	"encoding/hex"
	"math/big"
	"regexp"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLastFinalBlock(t *testing.T) {

	getBlockGetter := func(expectedBlockNum uint64) func(uint64) *types.Block {
		return func(in uint64) *types.Block {
			if in != expectedBlockNum {
				panic("unexpected block number")
			}
			header := &types.Header{Number: big.NewInt(int64(expectedBlockNum))}
			return types.NewBlockWithHeader(header)
		}
	}

	tests := []struct {
		name         string
		current      *types.Block
		candidate    *types.Header
		threshold    uint64
		getBlockNum  func(uint64) *types.Block
		reprocessing bool
		expect       uint64
		expectNil    bool
	}{
		{
			name:      "candidate is good",
			current:   types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}),
			candidate: &types.Header{Number: big.NewInt(5)},
			expect:    5,
		},
		{
			name:      "candidate is good, with forceThreshold",
			current:   types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}),
			candidate: &types.Header{Number: big.NewInt(5)},
			threshold: 200,
			expect:    5,
		},
		{
			name:         "is reprocessing",
			current:      types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}),
			candidate:    &types.Header{Number: big.NewInt(5)},
			reprocessing: true,
			expect:       10,
		},
		{
			name:      "no candidate, no threshold",
			current:   types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}),
			expectNil: true,
		},
		{
			name:      "no candidate, below threshold",
			current:   types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}),
			threshold: 200,
			expectNil: true,
		},
		{
			name:        "no candidate, above threshold",
			current:     types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1000)}),
			threshold:   200,
			getBlockNum: getBlockGetter(800),
			expect:      800,
		},
		{
			name:        "candidate too old",
			current:     types.NewBlockWithHeader(&types.Header{Number: big.NewInt(1000)}),
			candidate:   &types.Header{Number: big.NewInt(5)},
			threshold:   200,
			getBlockNum: getBlockGetter(800),
			expect:      800,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReprocessingWithSyncTarget = tt.reprocessing
			ForceFinalizedBlockAboveThreshold = tt.threshold

			if tt.expectNil {
				assert.Nil(t, LastFinalBlock(tt.current, tt.candidate, tt.getBlockNum))
			} else {
				lib := LastFinalBlock(tt.current, tt.candidate, tt.getBlockNum)
				require.NotNil(t, lib)
				assert.Equal(t, tt.expect, lib.Number.Uint64())
			}
		})
	}
}

func TestAccessList_marshal(t *testing.T) {
	tests := []struct {
		name    string
		l       AccessList
		wantOut string
	}{
		{"empty", nil, "00"},

		{
			"one address no keys",
			AccessList{
				types.AccessTuple{
					Address: address(t, "0x1234567890123456789012345678901234567890"),
				},
			},
			"01123456789012345678901234567890123456789000",
		},

		{
			"one address one key",
			AccessList{
				types.AccessTuple{
					Address:     address(t, "0x1234567890123456789012345678901234567890"),
					StorageKeys: []common.Hash{hash(t, "AB")},
				},
			},
			"0112345678901234567890123456789012345678900100000000000000000000000000000000000000000000000000000000000000ab",
		},

		{
			"one address multi keys",
			AccessList{
				types.AccessTuple{
					Address:     address(t, "0x1234567890123456789012345678901234567890"),
					StorageKeys: []common.Hash{hash(t, "AB"), hash(t, "EF")},
				},
			},
			`
				01
				1234567890123456789012345678901234567890
				  02
				  00000000000000000000000000000000000000000000000000000000000000ab
				  00000000000000000000000000000000000000000000000000000000000000ef
		   `,
		},

		{
			"multi address multi keys",
			AccessList{
				types.AccessTuple{
					Address:     address(t, "0x1234567890123456789012345678901234567890"),
					StorageKeys: []common.Hash{hash(t, "AB"), hash(t, "EF")},
				},
				types.AccessTuple{
					Address: address(t, "0xabcdefabcdefabcdefabcdefabcdefabcdef0910"),
				},
				types.AccessTuple{
					Address:     address(t, "0x1234567890123456789012345678901234567890"),
					StorageKeys: []common.Hash{hash(t, "12")},
				},
			},
			`
				03
				1234567890123456789012345678901234567890
					02
					00000000000000000000000000000000000000000000000000000000000000ab
					00000000000000000000000000000000000000000000000000000000000000ef
				abcdefabcdefabcdefabcdefabcdefabcdef0910
					00
				1234567890123456789012345678901234567890
					01
					0000000000000000000000000000000000000000000000000000000000000012
		   `,
		},
	}

	spaceRegex := regexp.MustCompile(`[\s\n\t]+`)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, spaceRegex.ReplaceAllString(tt.wantOut, ""), hex.EncodeToString(tt.l.marshal()))
		})
	}
}

func address(t *testing.T, in string) common.Address {
	t.Helper()

	addr, err := common.NewMixedcaseAddressFromString(in)
	require.NoError(t, err)

	return addr.Address()
}

func hash(t *testing.T, in string) common.Hash {
	t.Helper()

	return common.HexToHash(in)
}

func TestEndBlock_DataMarshaling(t *testing.T) {
	header := &types.Header{}
	uncles := []*types.Block{}
	finalizedBlock := types.NewBlockWithHeader(&types.Header{
		Number: big.NewInt(2048),
	})

	endData := map[string]interface{}{
		"header":             header,
		"uncles":             uncles,
		"totalDifficulty":    (*hexutil.Big)(big.NewInt(1024)),
		"finalizedBlockNum":  (*hexutil.Big)(finalizedBlock.Header().Number),
		"finalizedBlockHash": finalizedBlock.Header().Hash(),
	}

	require.Equal(t, `{"finalizedBlockHash":"0x38b1c79e3acb45df1ac1fbd4d70e08c655874fc592c5c86342a29baf4f001769","finalizedBlockNum":"0x800","header":{"parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000","sha3Uncles":"0x0000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","stateRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","transactionsRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","receiptsRoot":"0x0000000000000000000000000000000000000000000000000000000000000000","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","difficulty":null,"number":null,"gasLimit":"0x0","gasUsed":"0x0","timestamp":"0x0","extraData":"0x","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","baseFeePerGas":null,"withdrawalsRoot":null,"blobGasUsed":null,"excessBlobGas":null,"hash":"0xc3bd2d00745c03048a5616146a96f5ff78e54efb9e5b04af208cdaff6f3830ee"},"totalDifficulty":"0x400","uncles":[]}`, JSON(endData))
}
