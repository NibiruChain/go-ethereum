package firehose

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// Enabled determines if firehose instrumentation is enabled. Controlling
// firehose behavior is then controlled via other flag like.
var Enabled = false

// SyncInstrumentationEnabled determines if standard block syncing prints to standard
// console or if it discards it entirely and does not print nothing.
//
// This feature is enabled by default, can be disabled on a miner node used for
// speculative execution.
var SyncInstrumentationEnabled = true

// MiningEnabled determines if mining code should stay enabled even when firehose
// is active. In normal production setup, we always activate firehose on syncer node
// only. However, on local development setup, one might need to test speculative execution
// code locally. To achieve this, it's possible to enable firehose on miner, disable
// sync instrumentation and enable mining. This way, new blocks are mined, sync logs are
// not printed and speculative execution log can be accumulated.
var MiningEnabled = false

// BlockProgressEnabled enable output of finalize block line only.
//
// Currently, when taking backups, the best way to know about current
// last block seen is to track firehose logging. However, while doing
// the actual backups, it's not necessary to print all firehose logs.
//
// This settings will only affect printing the finalized block log line
// and will not impact any other firehose logs. If you need firehose
// instrumentation, activate firehose. The firehose setting has
// precedence over this setting.
var BlockProgressEnabled = false

// GenesisConfig keeps globally for the process the genesis config of the chain.
// The genesis config extracted from the initialization code of Geth, otherwise
// the operator will need to set the flag `--firehose-genesis-file` pointing
// it to correct genesis.json file for the chain.
//
// **Note** We use `interface{}` here instead of `*core.Genesis` because we otherwise
// have a compilation cycle because `core` package already uses `firehose` package.
// Consumer of this library make the cast back to the correct types when needed.
var GenesisConfig interface{}

var MissingGenesisPanicMessage = "Firehose requires to have the genesis config to properly emit genesis block for this chain " +
	"but it appears it was not set properly. Ensure you are using either chain's specific flag like " +
	"'--mainnet' or if using a custom network, you can use '--firehose-genesis' flag to provide. Firehose " +
	"is going to validate it against what your Geth database contains, so can be sure that it's going to " +
	"match what the databse have."

// Init initializes firehose with the given parameters.
//
// We cannot depend on `core` package because it already depends on `firehose` package. That's why here you see `genesis interface{}`
// (which should have been `*core.Genesis`) and a provided way to decode the genesis reader into its correct type.
func Init(
	enabled bool,
	syncInstrumentation bool,
	miningEnabled bool,
	blockProgress bool,
	genesis interface{},
	genesisFile string,
	newGenesis func() interface{},
	gethVersion string,
) error {
	log.Debug("Initializing firehose")
	Enabled = enabled
	SyncInstrumentationEnabled = syncInstrumentation
	MiningEnabled = miningEnabled
	BlockProgressEnabled = blockProgress

	genesisProvenance := "unset"

	// We must check for both `nil` and `(*core.Genesis)(nil)`, latter case that is not catch by using `genesis == nil` directly
	if !isNilInterfaceOrNilValue(genesis) {
		GenesisConfig = genesis
		genesisProvenance = "Geth Specific Flag (--<chain>)"
	} else {
		if genesisFilePath := genesisFile; genesisFilePath != "" {
			file, err := os.Open(genesisFilePath)
			if err != nil {
				return fmt.Errorf("firehose open genesis file: %w", err)
			}
			defer file.Close()

			var genesis = newGenesis()
			if err := json.NewDecoder(file).Decode(genesis); err != nil {
				return fmt.Errorf("decode genesis file %q: %w", genesisFilePath, err)
			}

			GenesisConfig = genesis
			genesisProvenance = "Firehose Specific Flag (--firehose-genesis <file>)"
		}
	}

	if Enabled {
		AllocateBuffers()
	}

	if Enabled || SyncInstrumentationEnabled || BlockProgressEnabled || MiningEnabled {
		log.Info("Firehose initialized",
			"enabled", Enabled,
			"sync_instrumentation_enabled", SyncInstrumentationEnabled,
			"mining_enabled", MiningEnabled,
			"block_progress_enabled", BlockProgressEnabled,
			"genesis_configured", genesis != nil,
			"genesis_provenance", genesisProvenance,
			"firehose_version", params.FirehoseVersion(),
			"geth_version", gethVersion,
			"chain_variant", params.Variant,
		)
	}

	MaybeSyncContext().InitVersion(
		gethVersion,
		params.FirehoseVersion(),
		params.Variant,
	)

	return nil
}

func isNilInterfaceOrNilValue(in interface{}) bool {
	if in == nil {
		return true
	}

	if rValue := reflect.ValueOf(in); rValue.Kind() == reflect.Ptr {
		return rValue.IsNil()
	}

	return false
}

// AllocateBuffers is called manually when Firehose is bootstrapped.
func AllocateBuffers() {
	if !Enabled {
		return
	}

	// 50 MiB
	BlockSyncBuffer = bytes.NewBuffer(make([]byte, 0, 50*1024*1024))

	// 5 MiB
	TxSyncBuffer = bytes.NewBuffer(make([]byte, 0, 5*1024*1024))
}

// BlockSyncBuffer to use and re-used for the state processor firehose context used to
// accumulate Firehose data for a block.
//
// BlockSyncBuffer is **not** thread-safe, it's expected to be used only by one thread at a time.
var BlockSyncBuffer *bytes.Buffer

// TxSyncBuffer holds a buffer of 5 MiB which should be enough for all transaction and it's
// re-used for all transactions so shouldn't be a big deal for the memory
//
// TxSyncBuffer is **not** thread-safe, it's expected to be used only by one thread at a time.
var TxSyncBuffer = bytes.NewBuffer(make([]byte, 0, 5*1024*1024))
