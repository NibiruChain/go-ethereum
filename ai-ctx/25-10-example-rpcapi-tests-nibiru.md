#  25-10-example-rpcapi-tests-nibiru.md

Reference E2E test outputs using a live network from Nibiru EVM. This shows geth
logging with `log.Printf` as well as other test logs.

=== RUN   Test
=== RUN   Test/TestExpectedMethods
=== RUN   Test/TestExpectedMethods/eth-rpcapi.EthAPI
=== RUN   Test/TestExpectedMethods/eth-rpcapi.FiltersAPI
=== RUN   Test/TestExpectedMethods/debug-rpcapi.DebugAPI
=== RUN   Test/TestParseBloomFromEvents
=== RUN   Test/TestParseBloomFromEvents/happy:_empty_events
=== RUN   Test/TestParseBloomFromEvents/happy:_events_with_bloom_included
=== NAME  Test
    cases.go:44: skipping integration test suite
--- SKIP: Test (0.03s)
    --- PASS: Test/TestExpectedMethods (0.00s)
        --- PASS: Test/TestExpectedMethods/eth-rpcapi.EthAPI (0.00s)
        --- PASS: Test/TestExpectedMethods/eth-rpcapi.FiltersAPI (0.00s)
        --- PASS: Test/TestExpectedMethods/debug-rpcapi.DebugAPI (0.00s)
    --- PASS: Test/TestParseBloomFromEvents (0.03s)
        --- PASS: Test/TestParseBloomFromEvents/happy:_empty_events (0.00s)
        --- PASS: Test/TestParseBloomFromEvents/happy:_events_with_bloom_included (0.03s)
=== RUN   TestBackendSuite
    backend_suite_test.go:65: ------------- SetupSuite: BEGIN ------------- 
    network.go:155: acquiring test network lock
    network.go:178: preparing test network...
    network.go:488: starting test network...
    start_node.go:135: Set EVM indexer
I[2025-10-15|15:48:41.899] service start                                indexer=evm msg="Starting EVMTxIndexerService service" impl=EVMTxIndexerService
E[2025-10-15|15:48:49.669] Tendermint WS rpcClient could not start      address=tcp://0.0.0.0:43103/websocket error="malformed HTTP response \"\\x00\\x00\\x06\\x04\\x00\\x00\\x00\\x00\\x00\\x00\\x05\\x00\\x00@\\x00\""
I[2025-10-15|15:48:49.672] Starting JSON-RPC server                     address=0.0.0.0:35659
I[2025-10-15|15:48:54.672] Starting JSON WebSocket server               address=127.0.0.1:8546
E[2025-10-15|15:48:54.673] Tendermint WS rpcClient could not start      address=tcp://0.0.0.0:43103/websocket error="malformed HTTP response \"\\x00\\x00\\x06\\x04\\x00\\x00\\x00\\x00\\x00\\x00\\x05\\x00\\x00@\\x00\""
    start_node.go:157: Set up Ethereum JSON-RPC client objects
    start_node.go:167: Expose typed methods for each namespace
    network.go:495: started validator 0
    network.go:504: started test network at height: 46
    backend_suite_test.go:81: SetupSuite: Funding `s.fundedAccEthAddr`: 0x44E2ADdd5685d680213D13440E0e5a40Ca140A43
    backend_suite_test.go:91: s.fundedEthAccAddr: 0x44E2ADdd5685d680213D13440E0e5a40Ca140A43, funds: 100000000unibi, s.node.Address: nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk
D[2025-10-15|15:48:56.694] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
I[2025-10-15|15:48:56.695] EthState AnteHandle BEGIN:
txhash: 0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:48:56.696] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:48:56.696] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:48:56.696] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:48:56.696] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:48:56.696] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:48:56.696] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:48:56.697] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:48:56.697] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 100000000000000000000
I[2025-10-15|15:48:56.697] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 99979000000000000000
I[2025-10-15|15:48:56.697] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:48:56.697] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:48:56.697] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:48:56.697] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:48:56.698] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:48:56.698] EthState AnteHandle END (SUCCESS):
txhash: 0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle BEGIN:
txhash: 0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 100000000000000000000
I[2025-10-15|15:48:57.226] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 99979000000000000000
I[2025-10-15|15:48:57.226] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:48:57.226] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:48:57.227] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:48:57.227] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:48:57.227] EthState AnteHandle END (SUCCESS):
txhash: 0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:48:57.227] UD-DEBUG EthereumTx START: hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09 module=geth 
I[2025-10-15|15:48:57.227] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:48:57.227] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:48:57.227] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:48:57.227] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:48:57.227] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:48:57.228] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:48:57.228] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:48:57.228] sdb.TxCfg(): {BlockHash:0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a TxHash:0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09 TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:48:57.228] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:48:57.228] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:48:57.228] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:48:57.228] UD-DEBUG EthereumTx TX SUCCESS: hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09, gasUsed=21000 module=geth 
    backend_suite_test.go:98: SetupSuite: Send Transfer TX and use the results in the tests
        transfer tx hash: 0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
D[2025-10-15|15:48:58.199] eth_getTransactionReceipt                    module=backend hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
    backend_suite_test.go:115: SetupSuite: Deploy test erc20 contract
D[2025-10-15|15:48:58.202] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5
I[2025-10-15|15:48:58.202] EthState AnteHandle BEGIN:
txhash: 0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:48:58.202] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:48:58.203] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99978999000000000000
I[2025-10-15|15:48:58.203] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 98478999000000000000
I[2025-10-15|15:48:58.204] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:48:58.204] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:48:58.204] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:48:58.204] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:48:58.204] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:48:58.204] EthState AnteHandle END (SUCCESS):
txhash: 0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle BEGIN:
txhash: 0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:48:58.277] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99978999000000000000
I[2025-10-15|15:48:58.278] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 98478999000000000000
I[2025-10-15|15:48:58.278] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:48:58.278] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:48:58.278] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:48:58.278] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:48:58.278] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:48:58.278] EthState AnteHandle END (SUCCESS):
txhash: 0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:48:58.279] UD-DEBUG EthereumTx START: hash=0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5 module=geth 
I[2025-10-15|15:48:58.279] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:48:58.279] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:48:58.279] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:48:58.279] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:48:58.279] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:48:58.281] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:48:58.281] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:48:58.281] sdb.TxCfg(): {BlockHash:0xc27d706d5d3230cf04f29d744eba3f74477a12219ceac790e6ee647a56414156 TxHash:0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5 TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:48:58.281] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:48:58.281] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:48:58.281] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:48:58.281] UD-DEBUG EthereumTx TX SUCCESS: hash=0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5, gasUsed=625412 module=geth 
D[2025-10-15|15:48:59.705] eth_getTransactionReceipt                    module=backend hash=0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5
    backend_suite_test.go:133: SuccessfulTx(transfer){ BlockNumber: 51, BlockHash: 0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a, TxHash: 0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09 }
    backend_suite_test.go:133: SuccessfulTx(deployContract){ BlockNumber: 53, BlockHash: 0xc27d706d5d3230cf04f29d744eba3f74477a12219ceac790e6ee647a56414156, TxHash: 0x6f64660d5b160e3b4684906b0d116b05986bfdcadb232bdd1ea8220f975324c5 }
    backend_suite_test.go:138: ------------- SetupSuite: END   ------------- 
=== RUN   TestBackendSuite/TestAccounts
=== RUN   TestBackendSuite/TestBlockNumber
=== RUN   TestBackendSuite/TestBlockNumberFromTendermint
=== RUN   TestBackendSuite/TestBlockNumberFromTendermint/happy:_block_number_specified
=== RUN   TestBackendSuite/TestBlockNumberFromTendermint/happy:_block_hash_specified
=== RUN   TestBackendSuite/TestBlockNumberFromTendermint/sad:_neither_block_number_nor_hash_specified
=== RUN   TestBackendSuite/TestChainConfig
=== RUN   TestBackendSuite/TestChainID
=== RUN   TestBackendSuite/TestCurrentHeader
=== RUN   TestBackendSuite/TestDoCall
    call_tx_test.go:87: txResponse from Backend.DoCall on penging block number: {"hash":"0x0000000000000000000000000000000000000000000000000000000000000000","logs":null,"gas_used":21000}
    call_tx_test.go:89: eth_call via RPC with the same block number should query latest and properly propagate the true error
D[2025-10-15|15:48:59.731] eth_call                                     client=json-rpc args="TransactionArgs{From:0x44E2ADdd5685d680213D13440E0e5a40Ca140A43, To:0xBff68D5AF95396168f20419FA02BF16a20BCD6bb, Gas:<nil>, Nonce:<nil>, Data:<nil>, Input:<nil>, AccessList:<nil>}" blocknumberorhash="{BlockNumber:0x4003fd4d50 BlockHash:<nil>}"
D[2025-10-15|15:48:59.732] eth_call failed                              client=json-rpc error="rpc error: code = Internal desc = insufficient balance for transfer"
I[2025-10-15|15:48:59.732] Served eth_call                              module=geth conn=127.0.0.1:43594 reqid=1 duration=1.821443ms err="rpc error: code = Internal desc = insufficient balance for transfer"
=== RUN   TestBackendSuite/TestEthBlockByNumber
=== RUN   TestBackendSuite/TestFeeHistory
=== RUN   TestBackendSuite/TestGasPrice
=== RUN   TestBackendSuite/TestGasUsedFunTokens
D[2025-10-15|15:49:00.773] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020
I[2025-10-15|15:49:00.774] EthState AnteHandle BEGIN:
txhash: 0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:00.774] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:00.774] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:00.774] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:00.774] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:00.774] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:00.774] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99353587000000000000
I[2025-10-15|15:49:00.775] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97853587000000000000
I[2025-10-15|15:49:00.775] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:00.775] EthState AnteHandle END (SUCCESS):
txhash: 0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
D[2025-10-15|15:49:00.776] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3
I[2025-10-15|15:49:00.776] EthState AnteHandle BEGIN:
txhash: 0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:00.776] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:00.776] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:00.776] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:00.776] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:00.776] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:00.776] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99353587000000000000
I[2025-10-15|15:49:00.777] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97853587000000000000
I[2025-10-15|15:49:00.777] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:00.777] EthState AnteHandle END (SUCCESS):
txhash: 0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
D[2025-10-15|15:49:00.778] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2
I[2025-10-15|15:49:00.778] EthState AnteHandle BEGIN:
txhash: 0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:00.778] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:00.778] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:00.778] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:00.778] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:00.778] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:00.778] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:00.779] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:00.779] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99353587000000000000
I[2025-10-15|15:49:00.779] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97853587000000000000
I[2025-10-15|15:49:00.779] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:00.779] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:00.779] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:00.779] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:00.779] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:00.780] EthState AnteHandle END (SUCCESS):
txhash: 0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle BEGIN:
txhash: 0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:01.068] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99353587000000000000
I[2025-10-15|15:49:01.069] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97853587000000000000
I[2025-10-15|15:49:01.069] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:01.069] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:01.069] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:01.069] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:01.069] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:01.069] EthState AnteHandle END (SUCCESS):
txhash: 0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:01.069] UD-DEBUG EthereumTx START: hash=0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020 module=geth 
I[2025-10-15|15:49:01.069] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:01.069] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:01.069] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:01.069] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:01.069] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:01.071] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:01.071] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:01.071] sdb.TxCfg(): {BlockHash:0xb5dded409716930a03dc479f64c348ae299c0f0bb83ea86c6df1ab5c87fc66bf TxHash:0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020 TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:49:01.071] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:01.071] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:01.071] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:01.071] UD-DEBUG EthereumTx TX SUCCESS: hash=0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020, gasUsed=30384 module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle BEGIN:
txhash: 0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:01.072] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 99323203000000000000
I[2025-10-15|15:49:01.073] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 30384000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97823203000000000000
I[2025-10-15|15:49:01.073] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1530384000000000000 module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:01.073] EthState AnteHandle END (SUCCESS):
txhash: 0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:01.073] UD-DEBUG EthereumTx START: hash=0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3 module=geth 
I[2025-10-15|15:49:01.073] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:01.073] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:01.073] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1530384000000000000 module=geth 
I[2025-10-15|15:49:01.073] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:01.073] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:01.074] sdb.TxCfg(): {BlockHash:0xb5dded409716930a03dc479f64c348ae299c0f0bb83ea86c6df1ab5c87fc66bf TxHash:0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3 TxIndex:1 LogIndex:7} module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx TX SUCCESS: hash=0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3, gasUsed=1500000 module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle BEGIN:
txhash: 0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97823203000000000000
I[2025-10-15|15:49:01.074] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1530384000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96323203000000000000
I[2025-10-15|15:49:01.074] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 3030384000000000000 module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:01.074] EthState AnteHandle END (SUCCESS):
txhash: 0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx START: hash=0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2 module=geth 
I[2025-10-15|15:49:01.074] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:01.075] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:01.075] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 3030384000000000000 module=geth 
I[2025-10-15|15:49:01.075] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:01.075] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:01.077] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:01.077] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:01.077] sdb.TxCfg(): {BlockHash:0xb5dded409716930a03dc479f64c348ae299c0f0bb83ea86c6df1ab5c87fc66bf TxHash:0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2 TxIndex:2 LogIndex:7} module=geth 
I[2025-10-15|15:49:01.077] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:01.077] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:01.077] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:01.077] UD-DEBUG EthereumTx TX SUCCESS: hash=0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2, gasUsed=30384 module=geth 
D[2025-10-15|15:49:01.280] eth_getTransactionReceipt                    module=backend hash=0x1211105bc75d82bc1aa89f690e6389a3285d9141d74040d6e548e55132252020
D[2025-10-15|15:49:01.783] eth_getTransactionReceipt                    module=backend hash=0x162488a55d82091663514ec9cb7a922f73ef781b01fac769535487142a86d4b3
D[2025-10-15|15:49:02.285] eth_getTransactionReceipt                    module=backend hash=0xefc01567e660d1952bd1c98085166c83c7e6b0e335a8cee074e6dd899640aaa2
=== RUN   TestBackendSuite/TestGasUsedTransfers
D[2025-10-15|15:49:03.289] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
I[2025-10-15|15:49:03.289] EthState AnteHandle BEGIN:
txhash: 0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:03.289] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:03.289] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:03.289] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:03.289] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97792819000000000000
I[2025-10-15|15:49:03.290] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97771819000000000000
I[2025-10-15|15:49:03.290] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:03.290] EthState AnteHandle END (SUCCESS):
txhash: 0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
D[2025-10-15|15:49:03.291] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236
I[2025-10-15|15:49:03.291] EthState AnteHandle BEGIN:
txhash: 0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97792819000000000000
I[2025-10-15|15:49:03.291] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97771819000000000000
I[2025-10-15|15:49:03.291] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:03.291] EthState AnteHandle END (SUCCESS):
txhash: 0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
D[2025-10-15|15:49:03.792] eth_getTransactionReceipt                    module=backend hash=0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
D[2025-10-15|15:49:03.792] tx not found                                 module=backend hash=0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b error="tx not found, hash: 0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b"
    backend_suite_test.go:249: Receipt still not available for tx 0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
I[2025-10-15|15:49:03.795] EthState AnteHandle BEGIN:
txhash: 0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:03.795] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97792819000000000000
I[2025-10-15|15:49:03.796] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97771819000000000000
I[2025-10-15|15:49:03.796] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:03.796] EthState AnteHandle END (SUCCESS):
txhash: 0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx START: hash=0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:03.796] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:03.796] sdb.TxCfg(): {BlockHash:0x0b59f6ea8e4f69a4a929d82183216ee432260a276a44b02418ed4b83ffdd1906 TxHash:0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:03.796] UD-DEBUG EthereumTx TX SUCCESS: hash=0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b, gasUsed=21000 module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle BEGIN:
txhash: 0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97771818000000000000
I[2025-10-15|15:49:03.797] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97750818000000000000
I[2025-10-15|15:49:03.797] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 42000000000000000 module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:03.797] EthState AnteHandle END (SUCCESS):
txhash: 0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx START: hash=0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236 module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:03.797] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 42000000000000000 module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:03.797] sdb.TxCfg(): {BlockHash:0x0b59f6ea8e4f69a4a929d82183216ee432260a276a44b02418ed4b83ffdd1906 TxHash:0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236 TxIndex:1 LogIndex:0} module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:03.797] UD-DEBUG EthereumTx TX SUCCESS: hash=0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236, gasUsed=21000 module=geth 
D[2025-10-15|15:49:04.292] eth_getTransactionReceipt                    module=backend hash=0x7f6b86a65b18082a7828d7e3125be1fe887b42551ed85e817345b5c3173ab71b
D[2025-10-15|15:49:04.792] eth_getTransactionReceipt                    module=backend hash=0x0cb78845d277a8d18051f9d7a1dbbfe1abbcd8b0c5ac961206fae51bcf06b236
=== RUN   TestBackendSuite/TestGetBalance
=== RUN   TestBackendSuite/TestGetBalance/happy:_funded_account_balance
=== RUN   TestBackendSuite/TestGetBalance/happy:_recipient_balance_at_block_1
=== RUN   TestBackendSuite/TestGetBalance/happy:_recipient_balance_after_transfer
=== RUN   TestBackendSuite/TestGetBalance/sad:_not_existing_account
=== RUN   TestBackendSuite/TestGetBlockByHash
D[2025-10-15|15:49:04.795] eth_getBlockByHash                           client=json-rpc hash=0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a fullTx=true
D[2025-10-15|15:49:04.796] Served eth_getBlockByHash                    module=geth conn=127.0.0.1:43594 reqid=2 duration=580.962┬╡s
    blocks_test.go:121: blockMap: map[baseFeePerGas:0xe8d4a51000 difficulty:0x0 extraData:0x gasLimit:0xffffffff gasUsed:0x5208 hash:0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a logsBloom:0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 miner:0xa320aed820f295642465bf6273d64c20371e5fe9 mixHash:0x0000000000000000000000000000000000000000000000000000000000000000 nonce:0x0000000000000000 number:0x33 parentHash:0xabfe6a70c3a4d8883a936418f0da6763cae4ca0655f3b36c94176f551dc9d1c1 receiptsRoot:0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421 sha3Uncles:0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347 size:0x3a7 stateRoot:0xf132e533abd3e9d2a63952bda9c9eeb01715187aa74eda7aab2da51e3a426be4 timestamp:0x68f008b8 totalDifficulty:0x0 transactions:[map[blockHash:0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a blockNumber:0x33 chainId:0x1b12 from:0x44e2addd5685d680213d13440e0e5a40ca140a43 gas:0x5208 gasPrice:0x1 hash:0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09 input:0x nonce:0x0 r:0x29727af49c8dadb351c9b8c2ac4c683ebb9813befc2323943cc1f056430ae6d0 s:0x647e3159e373e162f5731253222ef6fa5ad5fbc17b6231e389aa2d4bc72df888 to:0xbff68d5af95396168f20419fa02bf16a20bcd6bb transactionIndex:0x0 type:0x0 v:0x3648 value:0xe8d4a51000]] transactionsRoot:0x518e5f9433eb9fcab49ff398c8cf4968d9417776df4b7bcdd4622fdb85a85ef9 uncles:[]]
=== RUN   TestBackendSuite/TestGetBlockByNumberr
=== RUN   TestBackendSuite/TestGetBlockTransactionCountByHash
=== RUN   TestBackendSuite/TestGetBlockTransactionCountByNumber
=== RUN   TestBackendSuite/TestGetCode
=== RUN   TestBackendSuite/TestGetCode/happy:_valid_contract_address
=== RUN   TestBackendSuite/TestGetCode/sad:_not_a_contract_address
=== RUN   TestBackendSuite/TestGetHexProofs
=== RUN   TestBackendSuite/TestGetHexProofs/Case_no_proof_provided
=== RUN   TestBackendSuite/TestGetHexProofs/Case_no_proof_data_provided
=== RUN   TestBackendSuite/TestGetHexProofs/Case_valid_proof_provided
=== RUN   TestBackendSuite/TestGetLogsFromBlockResults
=== RUN   TestBackendSuite/TestGetProof
=== RUN   TestBackendSuite/TestGetProof/happy:_balance_of_the_contract_deployer
=== RUN   TestBackendSuite/TestGetProof/sad:_address_which_is_not_in_contract_storage
=== RUN   TestBackendSuite/TestGetStorageAt
=== RUN   TestBackendSuite/TestGetStorageAt/happy:_balance_of_the_contract_deployer
=== RUN   TestBackendSuite/TestGetStorageAt/sad:_address_which_is_not_in_contract_storage
=== RUN   TestBackendSuite/TestGetTransactionByBlockHashAndIndex
=== RUN   TestBackendSuite/TestGetTransactionByBlockHashAndIndex/happy:_tx_found
D[2025-10-15|15:49:04.799] eth_getTransactionByBlockHashAndIndex        client=json-rpc hash=0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a index=0x0
D[2025-10-15|15:49:04.799] Served eth_getTransactionByBlockHashAndIndex module=geth conn=127.0.0.1:43594 reqid=3 duration=250.7┬╡s
=== RUN   TestBackendSuite/TestGetTransactionByBlockHashAndIndex/sad:_block_not_found
D[2025-10-15|15:49:04.800] eth_getTransactionByBlockHashAndIndex        client=json-rpc hash=0x0000000000000000000000000000000000000000000000000000000000307830 index=0x1
D[2025-10-15|15:49:04.800] eth_getTransactionByBlockHashAndIndex failed client=json-rpc error="block not found: blockHash 0x0000000000000000000000000000000000000000000000000000000000307830: block query succeeded, but the block was nil"
I[2025-10-15|15:49:04.800] Served eth_getTransactionByBlockHashAndIndex module=geth conn=127.0.0.1:43594 reqid=4 duration=47.482┬╡s err="block not found: blockHash 0x0000000000000000000000000000000000000000000000000000000000307830: block query succeeded, but the block was nil"
=== RUN   TestBackendSuite/TestGetTransactionByBlockHashAndIndex/sad:_tx_not_found
D[2025-10-15|15:49:04.800] eth_getTransactionByBlockHashAndIndex        client=json-rpc hash=0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a index=0x270f
D[2025-10-15|15:49:04.800] block txs index out of bound                 module=backend index=9999
D[2025-10-15|15:49:04.800] Served eth_getTransactionByBlockHashAndIndex module=geth conn=127.0.0.1:43594 reqid=5 duration=122.004┬╡s
=== RUN   TestBackendSuite/TestGetTransactionByBlockNumberAndIndex
=== RUN   TestBackendSuite/TestGetTransactionByBlockNumberAndIndex/happy:_tx_found
D[2025-10-15|15:49:04.801] eth_getTransactionByBlockNumberAndIndex      client=json-rpc number=51 index=0x0
D[2025-10-15|15:49:04.801] Served eth_getTransactionByBlockNumberAndIndex module=geth conn=127.0.0.1:43594 reqid=6 duration=138.412┬╡s
=== RUN   TestBackendSuite/TestGetTransactionByBlockNumberAndIndex/sad:_block_not_found
D[2025-10-15|15:49:04.801] eth_getTransactionByBlockNumberAndIndex      client=json-rpc number=9999999 index=0x0
D[2025-10-15|15:49:04.801] eth_getTransactionByBlockNumberAndIndex failed client=json-rpc error="block not found: tendermint client failed to get block 9999999: height 9999999 must be less than or equal to the current blockchain height 64"
I[2025-10-15|15:49:04.801] Served eth_getTransactionByBlockNumberAndIndex module=geth conn=127.0.0.1:43594 reqid=7 duration=37.674┬╡s err="block not found: tendermint client failed to get block 9999999: height 9999999 must be less than or equal to the current blockchain height 64"
=== RUN   TestBackendSuite/TestGetTransactionByBlockNumberAndIndex/sad:_tx_not_found
D[2025-10-15|15:49:04.801] eth_getTransactionByBlockNumberAndIndex      client=json-rpc number=51 index=0x270f
D[2025-10-15|15:49:04.802] block txs index out of bound                 module=backend index=9999
D[2025-10-15|15:49:04.802] Served eth_getTransactionByBlockNumberAndIndex module=geth conn=127.0.0.1:43594 reqid=8 duration=250.425┬╡s
=== RUN   TestBackendSuite/TestGetTransactionByHash
=== RUN   TestBackendSuite/TestGetTransactionByHash/happy:_tx_found
D[2025-10-15|15:49:04.802] eth_getTransactionByHash                     client=json-rpc hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
D[2025-10-15|15:49:04.802] Served eth_getTransactionByHash              module=geth conn=127.0.0.1:43594 reqid=9 duration=132.179┬╡s
=== RUN   TestBackendSuite/TestGetTransactionByHash/sad:_tx_not_found
D[2025-10-15|15:49:04.802] eth_getTransactionByHash                     client=json-rpc hash=0x0000000000000000000000000000000000000000000000000000000000307830
D[2025-10-15|15:49:04.802] eth_getTransactionByHash failed              client=json-rpc error="no confirmed (pending) or unconfirmed tx found: tx not found, hash: 0x0000000000000000000000000000000000000000000000000000000000307830: no pending tx found with hash 0x0000000000000000000000000000000000000000000000000000000000307830"
I[2025-10-15|15:49:04.802] Served eth_getTransactionByHash              module=geth conn=127.0.0.1:43594 reqid=10 duration=31.166┬╡s err="no confirmed (pending) or unconfirmed tx found: tx not found, hash: 0x0000000000000000000000000000000000000000000000000000000000307830: no pending tx found with hash 0x0000000000000000000000000000000000000000000000000000000000307830"
=== RUN   TestBackendSuite/TestGetTransactionReceipt
=== RUN   TestBackendSuite/TestGetTransactionReceipt/happy:_tx_found
D[2025-10-15|15:49:04.803] eth_getTransactionReceipt                    client=json-rpc hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
D[2025-10-15|15:49:04.803] eth_getTransactionReceipt                    module=backend hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
D[2025-10-15|15:49:04.803] Served eth_getTransactionReceipt             module=geth conn=127.0.0.1:43594 reqid=11 duration=212.384┬╡s
=== RUN   TestBackendSuite/TestGetTransactionReceipt/sad:_tx_not_found
D[2025-10-15|15:49:04.804] eth_getTransactionReceipt                    client=json-rpc hash=0x0000000000000000000000000000000000000000000000000000000000307830
D[2025-10-15|15:49:04.804] eth_getTransactionReceipt                    module=backend hash=0x0000000000000000000000000000000000000000000000000000000000307830
D[2025-10-15|15:49:04.804] tx not found                                 module=backend hash=0x0000000000000000000000000000000000000000000000000000000000307830 error="tx not found, hash: 0x0000000000000000000000000000000000000000000000000000000000307830"
D[2025-10-15|15:49:04.804] Served eth_getTransactionReceipt             module=geth conn=127.0.0.1:43594 reqid=12 duration=28.599┬╡s
=== RUN   TestBackendSuite/TestGlobalMinGasPrice
=== RUN   TestBackendSuite/TestLogs
    tx_logs_test.go:43: TX1: Send simple nibi transfer
D[2025-10-15|15:49:05.807] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002
I[2025-10-15|15:49:05.807] EthState AnteHandle BEGIN:
txhash: 0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:05.807] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:05.807] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:05.807] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:05.807] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97750817000000000000
I[2025-10-15|15:49:05.808] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97729817000000000000
I[2025-10-15|15:49:05.808] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:05.808] EthState AnteHandle END (SUCCESS):
txhash: 0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
    tx_logs_test.go:50: TX2: Deploy ERC20 contract
D[2025-10-15|15:49:05.810] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a
I[2025-10-15|15:49:05.810] EthState AnteHandle BEGIN:
txhash: 0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:05.810] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97750817000000000000
I[2025-10-15|15:49:05.810] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96250817000000000000
I[2025-10-15|15:49:05.810] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:05.811] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:05.811] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:05.811] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:05.811] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:05.811] EthState AnteHandle END (SUCCESS):
txhash: 0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
    tx_logs_test.go:61: TX3: Create FunToken from ERC20
    tx_logs_test.go:76: TX4: Create FunToken from unibi coin
    tx_logs_test.go:98: TX5: Convert coin to EVM
    tx_logs_test.go:116: TX6: Send erc20 token to coin using precompile
D[2025-10-15|15:49:05.857] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1
I[2025-10-15|15:49:05.857] EthState AnteHandle BEGIN:
txhash: 0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97750817000000000000
I[2025-10-15|15:49:05.858] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 3000000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96250817000000000000
I[2025-10-15|15:49:05.858] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1503000000000000000 module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:05.858] EthState AnteHandle END (SUCCESS):
txhash: 0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
    tx_logs_test.go:137: Wait for all txs to be included in a block
I[2025-10-15|15:49:05.930] EthState AnteHandle BEGIN:
txhash: 0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:05.930] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:05.930] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:05.930] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:05.930] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:05.930] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:05.930] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97750817000000000000
I[2025-10-15|15:49:05.931] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 97729817000000000000
I[2025-10-15|15:49:05.931] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:05.931] EthState AnteHandle END (SUCCESS):
txhash: 0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:05.931] UD-DEBUG EthereumTx START: hash=0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002 module=geth 
I[2025-10-15|15:49:05.931] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:05.931] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:05.931] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:05.931] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:05.931] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:05.932] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:05.932] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:05.932] sdb.TxCfg(): {BlockHash:0x4140ca9fe1d600f0122c4f0a913b2ddbb4224fb328962fc73c9a06dbf084d2b7 TxHash:0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002 TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:49:05.932] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:05.932] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:05.932] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:05.933] UD-DEBUG EthereumTx TX SUCCESS: hash=0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002, gasUsed=21000 module=geth 
I[2025-10-15|15:49:05.942] EthState AnteHandle BEGIN:
txhash: 0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:05.943] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:05.944] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97729816000000000000
I[2025-10-15|15:49:05.944] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96229816000000000000
I[2025-10-15|15:49:05.944] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1521000000000000000 module=geth 
I[2025-10-15|15:49:05.944] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:05.944] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:05.944] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:05.944] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:05.944] EthState AnteHandle END (SUCCESS):
txhash: 0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:05.944] UD-DEBUG EthereumTx START: hash=0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a module=geth 
I[2025-10-15|15:49:05.944] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:05.944] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:05.944] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1521000000000000000 module=geth 
I[2025-10-15|15:49:05.944] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:05.944] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:05.945] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:05.945] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:05.945] sdb.TxCfg(): {BlockHash:0x4140ca9fe1d600f0122c4f0a913b2ddbb4224fb328962fc73c9a06dbf084d2b7 TxHash:0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a TxIndex:1 LogIndex:0} module=geth 
I[2025-10-15|15:49:05.945] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:05.945] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:05.945] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:05.945] UD-DEBUG EthereumTx TX SUCCESS: hash=0x58681724dfc68c0ba77e72a15ed8bcf07c55b9135219d63ebdbdc4c8434dd67a, gasUsed=625412 module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle BEGIN:
txhash: 0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97104404000000000000
I[2025-10-15|15:49:05.956] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 649412000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95604404000000000000
I[2025-10-15|15:49:05.956] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 2149412000000000000 module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:05.956] EthState AnteHandle END (SUCCESS):
txhash: 0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:05.957] UD-DEBUG EthereumTx START: hash=0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1 module=geth 
I[2025-10-15|15:49:05.957] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:05.957] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:05.957] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 2149412000000000000 module=geth 
I[2025-10-15|15:49:05.957] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:05.957] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:05.959] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:05.959] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:05.959] sdb.TxCfg(): {BlockHash:0x4140ca9fe1d600f0122c4f0a913b2ddbb4224fb328962fc73c9a06dbf084d2b7 TxHash:0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1 TxIndex:2 LogIndex:1} module=geth 
I[2025-10-15|15:49:05.959] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:05.959] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:05.959] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:05.960] UD-DEBUG EthereumTx TX SUCCESS: hash=0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1, gasUsed=30384 module=geth 
D[2025-10-15|15:49:06.358] eth_getTransactionReceipt                    module=backend hash=0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002
D[2025-10-15|15:49:06.859] eth_getTransactionReceipt                    module=backend hash=0x0885886f1655bbfd747347ebc00ea7b02301a404c7ecae7f6291d224314509c1
    tx_logs_test.go:151: txResp for tx3: {
          "data": "EnMKJS9ldGguZXZtLnYxLk1zZ0NyZWF0ZUZ1blRva2VuUmVzcG9uc2USSgpIChTJA1DcTZmMEVq6q/FojtyHLm/N/xIwZXJjMjAvMHhDOTAzNTBkYzREOTk4QzExNUFiYWFCZjE2ODhlZGM4NzJlNkZjZEZm",
          "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/eth.evm.v1.MsgCreateFunToken\"},{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"module\",\"value\":\"evm\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"wei_change\",\"attributes\":[{\"key\":\"wei_change_reason\",\"value\":\"bank.SendCoins\"},{\"key\":\"wei_change_addrs\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"burn\",\"attributes\":[{\"key\":\"burner\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"wei_change\",\"attributes\":[{\"key\":\"wei_change_reason\",\"value\":\"bank.BurnCoins\"},{\"key\":\"wei_change_addrs\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"}]},{\"type\":\"eth.evm.v1.EventFunTokenCreated\",\"attributes\":[{\"key\":\"bank_denom\",\"value\":\"\\\"erc20/0xC90350dc4D998C115AbaaBf1688edc872e6FcdFf\\\"\"},{\"key\":\"creator\",\"value\":\"\\\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\\\"\"},{\"key\":\"erc20_contract_address\",\"value\":\"\\\"0xC90350dc4D998C115AbaaBf1688edc872e6FcdFf\\\"\"},{\"key\":\"is_made_from_coin\",\"value\":\"false\"}]}]}]",
          "gas_wanted": 10000000,
          "gas_used": 356780,
          "events": [
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_received",
              "attributes": [
                {
                  "key": "receiver",
                  "value": "nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.SendCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "fee",
                  "value": "1000unibi",
                  "index": true
                },
                {
                  "key": "fee_payer",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "acc_seq",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk/3",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "signature",
                  "value": "oT1Js+lzoXMA4Umpc0RRSNUFtyWXTblxT7GPZWhKeP1RBFX8XLswWPbZOxY3SsOhbgwVA1HcJwwDUWxJAswbqQ==",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "action",
                  "value": "/eth.evm.v1.MsgCreateFunToken",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "module",
                  "value": "evm",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_received",
              "attributes": [
                {
                  "key": "receiver",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.SendCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "burn",
              "attributes": [
                {
                  "key": "burner",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.BurnCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                }
              ]
            },
            {
              "type": "eth.evm.v1.EventFunTokenCreated",
              "attributes": [
                {
                  "key": "bank_denom",
                  "value": "\"erc20/0xC90350dc4D998C115AbaaBf1688edc872e6FcdFf\"",
                  "index": true
                },
                {
                  "key": "creator",
                  "value": "\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"",
                  "index": true
                },
                {
                  "key": "erc20_contract_address",
                  "value": "\"0xC90350dc4D998C115AbaaBf1688edc872e6FcdFf\"",
                  "index": true
                },
                {
                  "key": "is_made_from_coin",
                  "value": "false",
                  "index": true
                }
              ]
            }
          ]
        }
    tx_logs_test.go:151: txResp for tx4: {
          "data": "EkoKJS9ldGguZXZtLnYxLk1zZ0NyZWF0ZUZ1blRva2VuUmVzcG9uc2USIQofChTNEKT97pzvt3MhYfSyCwGLo/Tn/xIFdW5pYmkYAQ==",
          "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/eth.evm.v1.MsgCreateFunToken\"},{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"module\",\"value\":\"evm\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"wei_change\",\"attributes\":[{\"key\":\"wei_change_reason\",\"value\":\"bank.SendCoins\"},{\"key\":\"wei_change_addrs\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"burn\",\"attributes\":[{\"key\":\"burner\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"},{\"key\":\"amount\",\"value\":\"10000000000unibi\"}]},{\"type\":\"wei_change\",\"attributes\":[{\"key\":\"wei_change_reason\",\"value\":\"bank.BurnCoins\"},{\"key\":\"wei_change_addrs\",\"value\":\"nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam\"}]},{\"type\":\"eth.evm.v1.EventTxLog\",\"attributes\":[{\"key\":\"logs\",\"value\":\"[{\\\"address\\\":\\\"0xcD10A4FdeE9CefB7732161f4B20b018bA3F4e7fF\\\",\\\"topics\\\":[\\\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\\\",\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"0x000000000000000000000000603871c2ddd41c26ee77495e2e31e6de7f9957e0\\\"],\\\"data\\\":null,\\\"block_number\\\":\\\"67\\\",\\\"tx_hash\\\":\\\"0x0000000000000000000000000000000000000000000000000000000000000000\\\",\\\"tx_index\\\":\\\"2\\\",\\\"block_hash\\\":\\\"0x4140ca9fe1d600f0122c4f0a913b2ddbb4224fb328962fc73c9a06dbf084d2b7\\\",\\\"index\\\":\\\"1\\\",\\\"removed\\\":false}]\"}]},{\"type\":\"eth.evm.v1.EventFunTokenCreated\",\"attributes\":[{\"key\":\"bank_denom\",\"value\":\"\\\"unibi\\\"\"},{\"key\":\"creator\",\"value\":\"\\\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\\\"\"},{\"key\":\"erc20_contract_address\",\"value\":\"\\\"0xcD10A4FdeE9CefB7732161f4B20b018bA3F4e7fF\\\"\"},{\"key\":\"is_made_from_coin\",\"value\":\"true\"}]}]}]",
          "gas_wanted": 10000000,
          "gas_used": 1518035,
          "events": [
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_received",
              "attributes": [
                {
                  "key": "receiver",
                  "value": "nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.SendCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "fee",
                  "value": "1000unibi",
                  "index": true
                },
                {
                  "key": "fee_payer",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "acc_seq",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk/4",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "signature",
                  "value": "Qv68/u73dc090/RnoYKyOwsZVr2MgBOlybof/0EA6kNGBiyQnjIRO4AMr+gh6f/7fW3SliLshX4kJx4W/FhBaw==",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "action",
                  "value": "/eth.evm.v1.MsgCreateFunToken",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "module",
                  "value": "evm",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_received",
              "attributes": [
                {
                  "key": "receiver",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.SendCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "burn",
              "attributes": [
                {
                  "key": "burner",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "10000000000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.BurnCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi1vqu8rska6swzdmnhf90zuv0xmelej4lq6hpdam",
                  "index": true
                }
              ]
            },
            {
              "type": "eth.evm.v1.EventTxLog",
              "attributes": [
                {
                  "key": "logs",
                  "value": "[{\"address\":\"0xcD10A4FdeE9CefB7732161f4B20b018bA3F4e7fF\",\"topics\":[\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\",\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"0x000000000000000000000000603871c2ddd41c26ee77495e2e31e6de7f9957e0\"],\"data\":null,\"block_number\":\"67\",\"tx_hash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"tx_index\":\"2\",\"block_hash\":\"0x4140ca9fe1d600f0122c4f0a913b2ddbb4224fb328962fc73c9a06dbf084d2b7\",\"index\":\"1\",\"removed\":false}]",
                  "index": true
                }
              ]
            },
            {
              "type": "eth.evm.v1.EventFunTokenCreated",
              "attributes": [
                {
                  "key": "bank_denom",
                  "value": "\"unibi\"",
                  "index": true
                },
                {
                  "key": "creator",
                  "value": "\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"",
                  "index": true
                },
                {
                  "key": "erc20_contract_address",
                  "value": "\"0xcD10A4FdeE9CefB7732161f4B20b018bA3F4e7fF\"",
                  "index": true
                },
                {
                  "key": "is_made_from_coin",
                  "value": "true",
                  "index": true
                }
              ]
            }
          ]
        }
    tx_logs_test.go:151: txResp for tx5: {
          "data": "EikKJy9ldGguZXZtLnYxLk1zZ0NvbnZlcnRDb2luVG9Fdm1SZXNwb25zZQ==",
          "log": "[{\"msg_index\":0,\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/eth.evm.v1.MsgConvertCoinToEvm\"},{\"key\":\"sender\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"},{\"key\":\"module\",\"value\":\"evm\"}]},{\"type\":\"wei_change\",\"attributes\":[{\"key\":\"wei_change_reason\",\"value\":\"evm.SubWei\"},{\"key\":\"wei_change_addrs\",\"value\":\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"}]},{\"type\":\"wei_change\",\"attributes\":[{\"key\":\"wei_change_reason\",\"value\":\"evm.AddWei\"},{\"key\":\"wei_change_addrs\",\"value\":\"nibi1pjk0v60cg347e2pxjyarc6uk4n2tq25hhymgg9\"}]},{\"type\":\"eth.evm.v1.EventConvertCoinToEvm\",\"attributes\":[{\"key\":\"bank_coin\",\"value\":\"{\\\"denom\\\":\\\"unibi\\\",\\\"amount\\\":\\\"1\\\"}\"},{\"key\":\"erc20_contract_address\",\"value\":\"\\\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\\\"\"},{\"key\":\"evm_logs\",\"value\":\"[{\\\"address\\\":\\\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\\\",\\\"topics\\\":[\\\"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c\\\",\\\"0x000000000000000000000000a320aed820f295642465bf6273d64c20371e5fe9\\\"],\\\"data\\\":\\\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAA=\\\"},{\\\"address\\\":\\\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\\\",\\\"topics\\\":[\\\"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c\\\",\\\"0x000000000000000000000000a320aed820f295642465bf6273d64c20371e5fe9\\\"],\\\"data\\\":\\\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAA=\\\"},{\\\"address\\\":\\\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\\\",\\\"topics\\\":[\\\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\\\",\\\"0x000000000000000000000000a320aed820f295642465bf6273d64c20371e5fe9\\\",\\\"0x00000000000000000000000044e2addd5685d680213d13440e0e5a40ca140a43\\\"],\\\"data\\\":\\\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAA=\\\"}]\"},{\"key\":\"sender\",\"value\":\"\\\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\\\"\"},{\"key\":\"to_eth_addr\",\"value\":\"\\\"0x44E2ADdd5685d680213D13440E0e5a40Ca140A43\\\"\"}]}]}]",
          "gas_wanted": 10000000,
          "gas_used": 700761,
          "events": [
            {
              "type": "coin_spent",
              "attributes": [
                {
                  "key": "spender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "coin_received",
              "attributes": [
                {
                  "key": "receiver",
                  "value": "nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "transfer",
              "attributes": [
                {
                  "key": "recipient",
                  "value": "nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "amount",
                  "value": "1000unibi",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "bank.SendCoins",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk, nibi17xpfvakm2amg962yls6f84z3kell8c5l8u8ezw",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "fee",
                  "value": "1000unibi",
                  "index": true
                },
                {
                  "key": "fee_payer",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "acc_seq",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk/5",
                  "index": true
                }
              ]
            },
            {
              "type": "tx",
              "attributes": [
                {
                  "key": "signature",
                  "value": "R237l4GGd4qbKwgK4+PSZ0zroUCeSlMd/r9Billz1TRoEGSbCSXx7ZPu7w0uSgpKtWXXfnMpXHH3/8gcwwmY+g==",
                  "index": true
                }
              ]
            },
            {
              "type": "message",
              "attributes": [
                {
                  "key": "action",
                  "value": "/eth.evm.v1.MsgConvertCoinToEvm",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                },
                {
                  "key": "module",
                  "value": "evm",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "evm.SubWei",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk",
                  "index": true
                }
              ]
            },
            {
              "type": "wei_change",
              "attributes": [
                {
                  "key": "wei_change_reason",
                  "value": "evm.AddWei",
                  "index": true
                },
                {
                  "key": "wei_change_addrs",
                  "value": "nibi1pjk0v60cg347e2pxjyarc6uk4n2tq25hhymgg9",
                  "index": true
                }
              ]
            },
            {
              "type": "eth.evm.v1.EventConvertCoinToEvm",
              "attributes": [
                {
                  "key": "bank_coin",
                  "value": "{\"denom\":\"unibi\",\"amount\":\"1\"}",
                  "index": true
                },
                {
                  "key": "erc20_contract_address",
                  "value": "\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\"",
                  "index": true
                },
                {
                  "key": "evm_logs",
                  "value": "[{\"address\":\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\",\"topics\":[\"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c\",\"0x000000000000000000000000a320aed820f295642465bf6273d64c20371e5fe9\"],\"data\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAA=\"},{\"address\":\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\",\"topics\":[\"0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c\",\"0x000000000000000000000000a320aed820f295642465bf6273d64c20371e5fe9\"],\"data\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAA=\"},{\"address\":\"0x0CaCF669f8446BeCA826913a3c6B96aCD4b02a97\",\"topics\":[\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\",\"0x000000000000000000000000a320aed820f295642465bf6273d64c20371e5fe9\",\"0x00000000000000000000000044e2addd5685d680213d13440e0e5a40ca140a43\"],\"data\":\"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6NSlEAA=\"}]",
                  "index": true
                },
                {
                  "key": "sender",
                  "value": "\"nibi15vs2akpq722kgfr9ha3884jvyqm3uhlfz7m6wk\"",
                  "index": true
                },
                {
                  "key": "to_eth_addr",
                  "value": "\"0x44E2ADdd5685d680213D13440E0e5a40Ca140A43\"",
                  "index": true
                }
              ]
            }
          ]
        }
    tx_logs_test.go:155: parse "eth.evm.v1.EventFunTokenCreated" from TX3
    tx_logs_test.go:171: parse "eth.evm.v1.EventFunTokenCreated" from TX4
    tx_logs_test.go:186: debugLogs: {
          "addr of recipient from tx 1": "0x4945c5A6C69917D286D3b9a8b1b5060c2877CDaF",
          "erc20 addr deployed in tx 2": "0xC90350dc4D998C115AbaaBf1688edc872e6FcdFf",
          "erc20FromCoinAddr (assumed funtoken address)": "0x66f67b74410a5c3493A8784108e2BbCC91F02DF7",
          "evm.EVM_MODULE_ADDRESS": "0x603871c2ddd41c26Ee77495E2E31e6De7f9957e0",
          "evm.FEE_COLLECTOR_ADDR": "0xf1829676DB577682E944fc3493d451B67Ff3E29F",
          "s.fundedAccEthAddr": "0x44E2ADdd5685d680213D13440E0e5a40Ca140A43",
          "s.node.EthAddress": "0xa320AeD820f295642465Bf6273D64c20371E5fE9",
          "tx hash of tx 1": "0x8af6e2295c968c6ec4fe3268bb585700b6fade5d59c5121f5b6c03f2a2695002"
        }
=== RUN   TestBackendSuite/TestLogs/TX1_-_simple_eth_transfer,_should_have_empty_logs
=== RUN   TestBackendSuite/TestLogs/TX2_-_deploying_erc20_contract,_should_have_logs
=== RUN   TestBackendSuite/TestLogs/TX3_-_create_FunToken_from_ERC20,_no_eth_tx,_no_logs
=== RUN   TestBackendSuite/TestLogs/TX4_-_create_FunToken_from_bank_coin,_no_eth_tx,_logs_for_contract_deployment
=== RUN   TestBackendSuite/TestLogs/TX5_-_Convert_coin_to_EVM,_no_eth_tx,_logs_for_minting_tokens_to_the_account
=== RUN   TestBackendSuite/TestLogs/TX6_-_Send_erc20_token_to_coin_using_precompile,_eth_tx,_logs_for_transferring_tokens_to_evm_module
=== RUN   TestBackendSuite/TestMultipleMsgsTxGasUsage
I[2025-10-15|15:49:06.864] EthState AnteHandle BEGIN:
txhash: 0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97074020000000000000
I[2025-10-15|15:49:06.864] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95574020000000000000
I[2025-10-15|15:49:06.864] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle END (SUCCESS):
txhash: 0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle BEGIN:
txhash: 0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97074020000000000000
I[2025-10-15|15:49:06.864] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96974020000000000000
I[2025-10-15|15:49:06.864] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 100000000000000000 module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:06.864] EthState AnteHandle END (SUCCESS):
txhash: 0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle BEGIN:
txhash: 0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97074020000000000000
I[2025-10-15|15:49:06.865] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96974020000000000000
I[2025-10-15|15:49:06.865] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 100000000000000000 module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:06.865] EthState AnteHandle END (SUCCESS):
txhash: 0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle BEGIN:
txhash: 0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 97074020000000000000
I[2025-10-15|15:49:07.033] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95574020000000000000
I[2025-10-15|15:49:07.033] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:07.033] EthState AnteHandle END (SUCCESS):
txhash: 0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:07.033] UD-DEBUG EthereumTx START: hash=0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f module=geth 
I[2025-10-15|15:49:07.033] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:07.033] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:07.033] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:07.033] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:07.033] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:07.034] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:07.034] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:07.034] sdb.TxCfg(): {BlockHash:0xc589895df8c04f889ef4becf546dfb0eafc5764b61886086cad5e3c3fc1ecde5 TxHash:0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:49:07.034] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:07.034] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:07.034] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:07.034] UD-DEBUG EthereumTx TX SUCCESS: hash=0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f, gasUsed=625412 module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle BEGIN:
txhash: 0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 96448608000000000000
I[2025-10-15|15:49:07.036] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 625412000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96348608000000000000
I[2025-10-15|15:49:07.036] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 725412000000000000 module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:07.036] EthState AnteHandle END (SUCCESS):
txhash: 0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:07.036] UD-DEBUG EthereumTx START: hash=0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b module=geth 
I[2025-10-15|15:49:07.036] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:07.036] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:07.036] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 725412000000000000 module=geth 
I[2025-10-15|15:49:07.036] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:07.036] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:07.037] sdb.TxCfg(): {BlockHash:0xc589895df8c04f889ef4becf546dfb0eafc5764b61886086cad5e3c3fc1ecde5 TxHash:0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b TxIndex:1 LogIndex:1} module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx TX SUCCESS: hash=0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b, gasUsed=51238 module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle BEGIN:
txhash: 0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 96397370000000000000
I[2025-10-15|15:49:07.037] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 676650000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96297370000000000000
I[2025-10-15|15:49:07.037] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 776650000000000000 module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:07.037] EthState AnteHandle END (SUCCESS):
txhash: 0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx START: hash=0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7 module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:07.037] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 776650000000000000 module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:07.037] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:07.038] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:07.038] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:07.038] sdb.TxCfg(): {BlockHash:0xc589895df8c04f889ef4becf546dfb0eafc5764b61886086cad5e3c3fc1ecde5 TxHash:0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7 TxIndex:2 LogIndex:2} module=geth 
I[2025-10-15|15:49:07.038] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:07.038] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:07.038] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:07.038] UD-DEBUG EthereumTx TX SUCCESS: hash=0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7, gasUsed=34138 module=geth 
D[2025-10-15|15:49:07.365] eth_getTransactionReceipt                    module=backend hash=0xe74d8f1e2c1b4cf09c15b07e51e6f8c9290c917b28d495b2c1b82f1c31e9df3f
D[2025-10-15|15:49:07.866] eth_getTransactionReceipt                    module=backend hash=0x759702db503749fd73d8ba6052f8a4af1515c533775360c250d9b231a2fe446b
D[2025-10-15|15:49:08.367] eth_getTransactionReceipt                    module=backend hash=0x672bd7c261d51ed79265549521b70debc3180dc9c26ee447cba2dbea037cd6a7
=== RUN   TestBackendSuite/TestNonceIncrementWithMultipleMsgsTx
    nonce_test.go:23: Before txs, nonce = 13
    nonce_test.go:43: Broadcast transaction. Expect failure in ante handler
    nonce_test.go:55: Nonce should be the same due to failure. Nonce only increase after successful txs.
    nonce_test.go:59: After failed txs, nonce = 13 (unchanged)
D[2025-10-15|15:49:09.380] eth_getTransactionReceipt                    module=backend hash=0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa
D[2025-10-15|15:49:09.380] tx not found                                 module=backend hash=0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa error="tx not found, hash: 0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa"
D[2025-10-15|15:49:09.380] eth_getTransactionReceipt                    module=backend hash=0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86
D[2025-10-15|15:49:09.380] tx not found                                 module=backend hash=0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86 error="tx not found, hash: 0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86"
D[2025-10-15|15:49:09.380] eth_getTransactionReceipt                    module=backend hash=0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a
D[2025-10-15|15:49:09.380] tx not found                                 module=backend hash=0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a error="tx not found, hash: 0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a"
    nonce_test.go:67: Broadcast 3 happy txs. Expect nonce to increment by 3
I[2025-10-15|15:49:09.381] EthState AnteHandle BEGIN:
txhash: 0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:09.381] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:09.381] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:09.381] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:09.381] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:09.381] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:09.381] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:09.382] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:09.382] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 96363232000000000000
I[2025-10-15|15:49:09.382] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 94863232000000000000
I[2025-10-15|15:49:09.383] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:09.383] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:09.383] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:09.383] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:09.383] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:09.383] EthState AnteHandle END (SUCCESS):
txhash: 0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
    nonce_test.go:76: sdk.TxResp creationTx: {
          "txhash": "0EA514F055F1C00EE9ABDE61C2C0DA09A7C480B30FDB0EC1EE21FCF9102E914A",
          "raw_log": "[]",
          "logs": [],
          "events": null
        }
I[2025-10-15|15:49:09.385] EthState AnteHandle BEGIN:
txhash: 0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:09.385] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:09.385] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:09.385] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:09.385] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:09.385] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:09.385] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:09.386] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:09.386] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 96363232000000000000
I[2025-10-15|15:49:09.386] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96263232000000000000
I[2025-10-15|15:49:09.386] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 100000000000000000 module=geth 
I[2025-10-15|15:49:09.386] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:09.386] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:09.386] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:09.387] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:09.387] EthState AnteHandle END (SUCCESS):
txhash: 0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
    nonce_test.go:76: sdk.TxResp firstTransferTx: {
          "txhash": "F2CA6677D6A9EDAF2A49B4A61351080459EAE1B1562283F5F3485A55AE3E6809",
          "raw_log": "[]",
          "logs": [],
          "events": null
        }
I[2025-10-15|15:49:09.390] EthState AnteHandle BEGIN:
txhash: 0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:09.390] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 96363232000000000000
I[2025-10-15|15:49:09.391] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 96263232000000000000
I[2025-10-15|15:49:09.391] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 100000000000000000 module=geth 
I[2025-10-15|15:49:09.391] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:09.391] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:09.391] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:09.391] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:09.391] EthState AnteHandle END (SUCCESS):
txhash: 0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
    nonce_test.go:76: sdk.TxResp secondTransferTx: {
          "txhash": "A52887D5ED3943806CC813E10923A5CBEE13EDED9D189BBB2D509208213D8F3E",
          "raw_log": "[]",
          "logs": [],
          "events": null
        }
I[2025-10-15|15:49:09.710] EthState AnteHandle BEGIN:
txhash: 0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:09.710] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 96363232000000000000
I[2025-10-15|15:49:09.711] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 94863232000000000000
I[2025-10-15|15:49:09.711] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:09.711] EthState AnteHandle END (SUCCESS):
txhash: 0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:09.711] UD-DEBUG EthereumTx START: hash=0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa module=geth 
I[2025-10-15|15:49:09.711] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:09.711] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:09.711] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 1500000000000000000 module=geth 
I[2025-10-15|15:49:09.711] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:09.711] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:09.712] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:09.712] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:09.712] sdb.TxCfg(): {BlockHash:0x0a1a24eeed6915aa82a76723b3d9b8237de2b047ca0ef88018616ab43b6b8feb TxHash:0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa TxIndex:0 LogIndex:0} module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx TX SUCCESS: hash=0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa, gasUsed=625412 module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle BEGIN:
txhash: 0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 95737820000000000000
I[2025-10-15|15:49:09.713] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 625412000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95637820000000000000
I[2025-10-15|15:49:09.713] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 725412000000000000 module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:09.713] EthState AnteHandle END (SUCCESS):
txhash: 0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx START: hash=0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86 module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:09.713] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:09.713] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 725412000000000000 module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:09.714] sdb.TxCfg(): {BlockHash:0x0a1a24eeed6915aa82a76723b3d9b8237de2b047ca0ef88018616ab43b6b8feb TxHash:0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86 TxIndex:1 LogIndex:1} module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:09.714] UD-DEBUG EthereumTx TX SUCCESS: hash=0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86, gasUsed=34138 module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle BEGIN:
txhash: 0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a
{ IsCheckTx false, IsDeliverTx true  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:09.715] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 95703682000000000000
I[2025-10-15|15:49:09.715] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 659550000000000000 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95603682000000000000
I[2025-10-15|15:49:09.716] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 759550000000000000 module=geth 
I[2025-10-15|15:49:09.716] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:09.716] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:09.716] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:09.716] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:09.716] EthState AnteHandle END (SUCCESS):
txhash: 0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a
{ IsCheckTx false, ReCheckTx false, IsDeliverTx true } module=geth 
I[2025-10-15|15:49:09.716] UD-DEBUG EthereumTx START: hash=0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a module=geth 
I[2025-10-15|15:49:09.716] UD-DEBUG EthereumTx ValidateBasic PASSED     module=geth 
I[2025-10-15|15:49:09.716] UD-DEBUG EthereumTx txConfig and evmCfg created module=geth 
I[2025-10-15|15:49:09.716] sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 759550000000000000 module=geth 
I[2025-10-15|15:49:09.716] UD-DEBUG EthereumTx TransactionToMessage PASSED module=geth 
I[2025-10-15|15:49:09.716] UD-DEBUG EthereumTx SDB and EVM created      module=geth 
I[2025-10-15|15:49:09.717] UD-DEBUG EthereumTx ApplyEvmMsg SUCCESS      module=geth 
I[2025-10-15|15:49:09.717] UD-DEBUG EthereumTx GAS CONSUMPTION SUCCESS  module=geth 
I[2025-10-15|15:49:09.717] sdb.TxCfg(): {BlockHash:0x0a1a24eeed6915aa82a76723b3d9b8237de2b047ca0ef88018616ab43b6b8feb TxHash:0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a TxIndex:2 LogIndex:2} module=geth 
I[2025-10-15|15:49:09.717] UD-DEBUG EthereumTx GAS REFUND SUCCESS       module=geth 
I[2025-10-15|15:49:09.717] UD-DEBUG EthereumTx EMIT EVENTS SUCCESS      module=geth 
I[2025-10-15|15:49:09.717] UD-DEBUG EthereumTx EMIT TX LOG SUCCESS      module=geth 
I[2025-10-15|15:49:09.718] UD-DEBUG EthereumTx TX SUCCESS: hash=0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a, gasUsed=34138 module=geth 
    nonce_test.go:84: Assert all transactions included in block
D[2025-10-15|15:49:10.894] eth_getTransactionReceipt                    module=backend hash=0x6919020fd565ac4bee74fa1de2c9472a304c8c10aa6304a2275740833f2a5aaa
D[2025-10-15|15:49:11.396] eth_getTransactionReceipt                    module=backend hash=0xce2779401af2e6c1c00c225f3479881b76f1330583cb9e9e17de9f5d9cff7d86
D[2025-10-15|15:49:11.897] eth_getTransactionReceipt                    module=backend hash=0xe9e16be44dde870c3e7c366ec9573b456452a81a797cf4bf258a5ca085ed7f3a
=== RUN   TestBackendSuite/TestPendingTransactions
D[2025-10-15|15:49:11.900] eth_sendRawTransaction beforeBroadcast       module=backend txHash=0x293b39b9e7d056a9da319e1d61698c2fc8284f2f43c2a7cead55ce8613234ea8
I[2025-10-15|15:49:11.900] EthState AnteHandle BEGIN:
txhash: 0x293b39b9e7d056a9da319e1d61698c2fc8284f2f43c2a7cead55ce8613234ea8
{ IsCheckTx true, IsDeliverTx false  ReCheckTxfalse } module=geth 
I[2025-10-15|15:49:11.900] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:11.900] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:11.900] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:11.900] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:11.900] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:11.900] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 95669544000000000000
I[2025-10-15|15:49:11.901] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95648544000000000000
I[2025-10-15|15:49:11.901] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:11.901] EthState AnteHandle END (SUCCESS):
txhash: 0x293b39b9e7d056a9da319e1d61698c2fc8284f2f43c2a7cead55ce8613234ea8
{ IsCheckTx true, ReCheckTx false, IsDeliverTx false } module=geth 
=== RUN   TestBackendSuite/TestRPCBlockRangeCap
=== RUN   TestBackendSuite/TestRPCEVMTimeout
=== RUN   TestBackendSuite/TestRPCFilterCap
=== RUN   TestBackendSuite/TestRPCGasCap
=== RUN   TestBackendSuite/TestRPCLogsCap
=== RUN   TestBackendSuite/TestRPCMinGasPrice
=== RUN   TestBackendSuite/TestReceiptMarshalJson
=== RUN   TestBackendSuite/TestSetTxDefaults
=== RUN   TestBackendSuite/TestSetTxDefaults/happy:_minimal_args_set
D[2025-10-15|15:49:11.921] estimate gas usage automatically             module=backend gas=0x5208
=== RUN   TestBackendSuite/TestSetTxDefaults/happy:_gas_price_set
I[2025-10-15|15:49:11.924] EthState AnteHandle BEGIN:
txhash: 0x293b39b9e7d056a9da319e1d61698c2fc8284f2f43c2a7cead55ce8613234ea8
{ IsCheckTx true, IsDeliverTx false  ReCheckTxtrue } module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 0 passed       module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 1 passed       module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 2 passed       module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 3 passed       module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 4 passed       module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 5 passed       module=geth 
I[2025-10-15|15:49:11.924] EthState AnteHandle Body elem 6 passed       module=geth 
I[2025-10-15|15:49:11.925] EthState AnteHandle Body elem 7 passed       module=geth 
EthAnteDeductGas Pre: sdb.GetBalance(msgEthTx.FromAddr()): 95669544000000000000
I[2025-10-15|15:49:11.925] EthAnteDeductGas Pre: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 0 module=geth 
EthAnteDeductGas Post: sdb.GetBalance(msgEthTx.FromAddr()): 95648544000000000000
I[2025-10-15|15:49:11.925] EthAnteDeductGas Post: sdb.GetBalance(evm.FEE_COLLECTOR_ADDR): 21000000000000000 module=geth 
I[2025-10-15|15:49:11.925] EthState AnteHandle Body elem 8 passed       module=geth 
I[2025-10-15|15:49:11.925] EthState AnteHandle Body elem 9 passed       module=geth 
I[2025-10-15|15:49:11.925] EthState AnteHandle Body elem 10 passed      module=geth 
I[2025-10-15|15:49:11.925] EthState AnteHandle Body elem 11 passed      module=geth 
I[2025-10-15|15:49:11.925] EthState AnteHandle END (SUCCESS):
txhash: 0x293b39b9e7d056a9da319e1d61698c2fc8284f2f43c2a7cead55ce8613234ea8
{ IsCheckTx true, ReCheckTx true, IsDeliverTx false } module=geth 
D[2025-10-15|15:49:11.951] estimate gas usage automatically             module=backend gas=0x5208
=== RUN   TestBackendSuite/TestSetTxDefaults/sad:_no_to_(contract_creation)_and_no_data
=== RUN   TestBackendSuite/TestSetTxDefaults/sad:_transfer_without_from_specified_generates_new_empty_account
=== RUN   TestBackendSuite/TestSuggestGasTipCap
=== RUN   TestBackendSuite/TestSyncing
=== RUN   TestBackendSuite/TestTraceBlock
=== RUN   TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_no_txs,_tracer:_default
D[2025-10-15|15:49:11.962] debug_traceBlockByNumber                     module=debug height=1
D[2025-10-15|15:49:11.962] Served debug_traceBlockByNumber              module=geth conn=127.0.0.1:43594 reqid=13 duration=157.936┬╡s
D[2025-10-15|15:49:11.962] debug_traceBlockByHash                       module=debug hash=0xd75e026d5019a4489c1a4450a3bd9f30af4a2434100e7c44a8ec57f71373adda
D[2025-10-15|15:49:11.962] Served debug_traceBlockByHash                module=geth conn=127.0.0.1:43594 reqid=14 duration=103.305┬╡s
    tracing_test.go:177: TraceBlock result: []
    tracing_test.go:177: TraceBlock result: []
    tracing_test.go:177: TraceBlock result: []
=== RUN   TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_no_txs,_tracer:_callTracer
D[2025-10-15|15:49:11.963] debug_traceBlockByNumber                     module=debug height=1
D[2025-10-15|15:49:11.963] Served debug_traceBlockByNumber              module=geth conn=127.0.0.1:43594 reqid=15 duration=82.497┬╡s
D[2025-10-15|15:49:11.963] debug_traceBlockByHash                       module=debug hash=0xd75e026d5019a4489c1a4450a3bd9f30af4a2434100e7c44a8ec57f71373adda
D[2025-10-15|15:49:11.963] Served debug_traceBlockByHash                module=geth conn=127.0.0.1:43594 reqid=16 duration=83.231┬╡s
    tracing_test.go:177: TraceBlock result: []
    tracing_test.go:177: TraceBlock result: []
    tracing_test.go:177: TraceBlock result: []
=== RUN   TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_transfer_tx,_tracer:_callTracer
D[2025-10-15|15:49:11.965] debug_traceBlockByNumber                     module=debug height=51
D[2025-10-15|15:49:11.965] Served debug_traceBlockByNumber              module=geth conn=127.0.0.1:43594 reqid=17 duration=717.083┬╡s
D[2025-10-15|15:49:11.966] debug_traceBlockByHash                       module=debug hash=0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a
D[2025-10-15|15:49:11.966] Served debug_traceBlockByHash                module=geth conn=127.0.0.1:43594 reqid=18 duration=674.367┬╡s
    tracing_test.go:177: TraceBlock result: [
          {
            "result": {
              "from": "0x44e2addd5685d680213d13440e0e5a40ca140a43",
              "gas": "0x5208",
              "gasUsed": "0x5208",
              "input": "0x",
              "to": "0xbff68d5af95396168f20419fa02bf16a20bcd6bb",
              "type": "CALL",
              "value": "0xe8d4a51000"
            }
          }
        ]
    tracing_test.go:177: TraceBlock result: [
          {
            "result": {
              "from": "0x44e2addd5685d680213d13440e0e5a40ca140a43",
              "gas": "0x5208",
              "gasUsed": "0x5208",
              "input": "0x",
              "to": "0xbff68d5af95396168f20419fa02bf16a20bcd6bb",
              "type": "CALL",
              "value": "0xe8d4a51000"
            }
          }
        ]
    tracing_test.go:177: TraceBlock result: [
          {
            "result": {
              "from": "0x44e2addd5685d680213d13440e0e5a40ca140a43",
              "gas": "0x5208",
              "gasUsed": "0x5208",
              "input": "0x",
              "to": "0xbff68d5af95396168f20419fa02bf16a20bcd6bb",
              "type": "CALL",
              "value": "0xe8d4a51000"
            }
          }
        ]
=== RUN   TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_transfer_tx,_tracer:_default
D[2025-10-15|15:49:11.967] debug_traceBlockByNumber                     module=debug height=51
D[2025-10-15|15:49:11.968] Served debug_traceBlockByNumber              module=geth conn=127.0.0.1:43594 reqid=19 duration=643.294┬╡s
D[2025-10-15|15:49:11.968] debug_traceBlockByHash                       module=debug hash=0x93beb013064c73737c05d3d33084dde98192a72fe43dac980b8371643ac4495a
D[2025-10-15|15:49:11.969] Served debug_traceBlockByHash                module=geth conn=127.0.0.1:43594 reqid=20 duration=783.997┬╡s
    tracing_test.go:177: TraceBlock result: [
          {
            "result": {
              "from": "0x44e2addd5685d680213d13440e0e5a40ca140a43",
              "gas": "0x5208",
              "gasUsed": "0x5208",
              "input": "0x",
              "to": "0xbff68d5af95396168f20419fa02bf16a20bcd6bb",
              "type": "CALL",
              "value": "0xe8d4a51000"
            }
          }
        ]
    tracing_test.go:177: TraceBlock result: [
          {
            "result": {
              "from": "0x44e2addd5685d680213d13440e0e5a40ca140a43",
              "gas": "0x5208",
              "gasUsed": "0x5208",
              "input": "0x",
              "to": "0xbff68d5af95396168f20419fa02bf16a20bcd6bb",
              "type": "CALL",
              "value": "0xe8d4a51000"
            }
          }
        ]
    tracing_test.go:177: TraceBlock result: [
          {
            "result": {
              "from": "0x44e2addd5685d680213d13440e0e5a40ca140a43",
              "gas": "0x5208",
              "gasUsed": "0x5208",
              "input": "0x",
              "to": "0xbff68d5af95396168f20419fa02bf16a20bcd6bb",
              "type": "CALL",
              "value": "0xe8d4a51000"
            }
          }
        ]
=== RUN   TestBackendSuite/TestTraceCall
E[2025-10-15|15:49:11.970] debug_traceCall                              module=debug err="unsupported key type"
D[2025-10-15|15:49:11.971] Served debug_traceCall                       module=geth conn=127.0.0.1:43594 reqid=21 duration=651.726┬╡s
=== RUN   TestBackendSuite/TestTraceTransaction
=== RUN   TestBackendSuite/TestTraceTransaction/sad:_tx_not_found
D[2025-10-15|15:49:11.971] tx not found                                 module=backend hash=0x0000000000000000000000000000000000000000000000000000000000307830
=== RUN   TestBackendSuite/TestTraceTransaction/happy:_tx_found
D[2025-10-15|15:49:11.972] debug_traceTransaction                       module=debug hash=0x055dafb97ba0a7331265d572c2a17ce59d6e08d06b4e6579300d009511389c09
D[2025-10-15|15:49:11.974] Served debug_traceTransaction                module=geth conn=127.0.0.1:43594 reqid=22 duration=1.777353ms
=== NAME  TestBackendSuite
    cases.go:76: Test suite succeeded on attempt 0
--- PASS: TestBackendSuite (42.43s)
    --- PASS: TestBackendSuite/TestAccounts (0.01s)
    --- PASS: TestBackendSuite/TestBlockNumber (0.00s)
    --- PASS: TestBackendSuite/TestBlockNumberFromTendermint (0.00s)
        --- PASS: TestBackendSuite/TestBlockNumberFromTendermint/happy:_block_number_specified (0.00s)
        --- PASS: TestBackendSuite/TestBlockNumberFromTendermint/happy:_block_hash_specified (0.00s)
        --- PASS: TestBackendSuite/TestBlockNumberFromTendermint/sad:_neither_block_number_nor_hash_specified (0.00s)
    --- PASS: TestBackendSuite/TestChainConfig (0.00s)
    --- PASS: TestBackendSuite/TestChainID (0.00s)
    --- PASS: TestBackendSuite/TestCurrentHeader (0.00s)
    --- PASS: TestBackendSuite/TestDoCall (0.01s)
    --- PASS: TestBackendSuite/TestEthBlockByNumber (0.00s)
    --- PASS: TestBackendSuite/TestFeeHistory (0.01s)
    --- PASS: TestBackendSuite/TestGasPrice (0.00s)
    --- PASS: TestBackendSuite/TestGasUsedFunTokens (2.54s)
    --- PASS: TestBackendSuite/TestGasUsedTransfers (2.51s)
    --- PASS: TestBackendSuite/TestGetBalance (0.00s)
        --- PASS: TestBackendSuite/TestGetBalance/happy:_funded_account_balance (0.00s)
        --- PASS: TestBackendSuite/TestGetBalance/happy:_recipient_balance_at_block_1 (0.00s)
        --- PASS: TestBackendSuite/TestGetBalance/happy:_recipient_balance_after_transfer (0.00s)
        --- PASS: TestBackendSuite/TestGetBalance/sad:_not_existing_account (0.00s)
    --- PASS: TestBackendSuite/TestGetBlockByHash (0.00s)
    --- PASS: TestBackendSuite/TestGetBlockByNumberr (0.00s)
    --- PASS: TestBackendSuite/TestGetBlockTransactionCountByHash (0.00s)
    --- PASS: TestBackendSuite/TestGetBlockTransactionCountByNumber (0.00s)
    --- PASS: TestBackendSuite/TestGetCode (0.00s)
        --- PASS: TestBackendSuite/TestGetCode/happy:_valid_contract_address (0.00s)
        --- PASS: TestBackendSuite/TestGetCode/sad:_not_a_contract_address (0.00s)
    --- PASS: TestBackendSuite/TestGetHexProofs (0.00s)
        --- PASS: TestBackendSuite/TestGetHexProofs/Case_no_proof_provided (0.00s)
        --- PASS: TestBackendSuite/TestGetHexProofs/Case_no_proof_data_provided (0.00s)
        --- PASS: TestBackendSuite/TestGetHexProofs/Case_valid_proof_provided (0.00s)
    --- PASS: TestBackendSuite/TestGetLogsFromBlockResults (0.00s)
    --- PASS: TestBackendSuite/TestGetProof (0.00s)
        --- PASS: TestBackendSuite/TestGetProof/happy:_balance_of_the_contract_deployer (0.00s)
        --- PASS: TestBackendSuite/TestGetProof/sad:_address_which_is_not_in_contract_storage (0.00s)
    --- PASS: TestBackendSuite/TestGetStorageAt (0.00s)
        --- PASS: TestBackendSuite/TestGetStorageAt/happy:_balance_of_the_contract_deployer (0.00s)
        --- PASS: TestBackendSuite/TestGetStorageAt/sad:_address_which_is_not_in_contract_storage (0.00s)
    --- PASS: TestBackendSuite/TestGetTransactionByBlockHashAndIndex (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByBlockHashAndIndex/happy:_tx_found (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByBlockHashAndIndex/sad:_block_not_found (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByBlockHashAndIndex/sad:_tx_not_found (0.00s)
    --- PASS: TestBackendSuite/TestGetTransactionByBlockNumberAndIndex (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByBlockNumberAndIndex/happy:_tx_found (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByBlockNumberAndIndex/sad:_block_not_found (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByBlockNumberAndIndex/sad:_tx_not_found (0.00s)
    --- PASS: TestBackendSuite/TestGetTransactionByHash (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByHash/happy:_tx_found (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionByHash/sad:_tx_not_found (0.00s)
    --- PASS: TestBackendSuite/TestGetTransactionReceipt (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionReceipt/happy:_tx_found (0.00s)
        --- PASS: TestBackendSuite/TestGetTransactionReceipt/sad:_tx_not_found (0.00s)
    --- PASS: TestBackendSuite/TestGlobalMinGasPrice (0.00s)
    --- PASS: TestBackendSuite/TestLogs (2.06s)
        --- PASS: TestBackendSuite/TestLogs/TX1_-_simple_eth_transfer,_should_have_empty_logs (0.00s)
        --- PASS: TestBackendSuite/TestLogs/TX2_-_deploying_erc20_contract,_should_have_logs (0.00s)
        --- PASS: TestBackendSuite/TestLogs/TX3_-_create_FunToken_from_ERC20,_no_eth_tx,_no_logs (0.00s)
        --- PASS: TestBackendSuite/TestLogs/TX4_-_create_FunToken_from_bank_coin,_no_eth_tx,_logs_for_contract_deployment (0.00s)
        --- PASS: TestBackendSuite/TestLogs/TX5_-_Convert_coin_to_EVM,_no_eth_tx,_logs_for_minting_tokens_to_the_account (0.00s)
        --- PASS: TestBackendSuite/TestLogs/TX6_-_Send_erc20_token_to_coin_using_precompile,_eth_tx,_logs_for_transferring_tokens_to_evm_module (0.00s)
    --- PASS: TestBackendSuite/TestMultipleMsgsTxGasUsage (1.51s)
    --- PASS: TestBackendSuite/TestNonceIncrementWithMultipleMsgsTx (3.52s)
    --- PASS: TestBackendSuite/TestPendingTransactions (0.00s)
    --- PASS: TestBackendSuite/TestRPCBlockRangeCap (0.00s)
    --- PASS: TestBackendSuite/TestRPCEVMTimeout (0.00s)
    --- PASS: TestBackendSuite/TestRPCFilterCap (0.00s)
    --- PASS: TestBackendSuite/TestRPCGasCap (0.00s)
    --- PASS: TestBackendSuite/TestRPCLogsCap (0.00s)
    --- PASS: TestBackendSuite/TestRPCMinGasPrice (0.00s)
    --- PASS: TestBackendSuite/TestReceiptMarshalJson (0.00s)
    --- PASS: TestBackendSuite/TestSetTxDefaults (0.05s)
        --- PASS: TestBackendSuite/TestSetTxDefaults/happy:_minimal_args_set (0.01s)
        --- PASS: TestBackendSuite/TestSetTxDefaults/happy:_gas_price_set (0.03s)
        --- PASS: TestBackendSuite/TestSetTxDefaults/sad:_no_to_(contract_creation)_and_no_data (0.00s)
        --- PASS: TestBackendSuite/TestSetTxDefaults/sad:_transfer_without_from_specified_generates_new_empty_account (0.01s)
    --- PASS: TestBackendSuite/TestSuggestGasTipCap (0.00s)
    --- PASS: TestBackendSuite/TestSyncing (0.00s)
    --- PASS: TestBackendSuite/TestTraceBlock (0.01s)
        --- PASS: TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_no_txs,_tracer:_default (0.00s)
        --- PASS: TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_no_txs,_tracer:_callTracer (0.00s)
        --- PASS: TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_transfer_tx,_tracer:_callTracer (0.00s)
        --- PASS: TestBackendSuite/TestTraceBlock/happy:_TraceBlock,_transfer_tx,_tracer:_default (0.00s)
    --- PASS: TestBackendSuite/TestTraceCall (0.00s)
    --- PASS: TestBackendSuite/TestTraceTransaction (0.00s)
        --- PASS: TestBackendSuite/TestTraceTransaction/sad:_tx_not_found (0.00s)
        --- PASS: TestBackendSuite/TestTraceTransaction/happy:_tx_found (0.00s)
=== RUN   TestEventSubscriber
E[2025-10-15|15:49:11.976] failed to add event topic to event bus       topic=event error="topic already registered"
--- PASS: TestEventSubscriber (0.00s)
PASS
ok  	github.com/NibiruChain/nibiru/v2/eth/rpc/rpcapi	42.520s

