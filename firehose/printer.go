package firehose

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Printer interface {
	// Write is a raw write to the printer mainly appending pre-formatted Firehose
	// lines in bytes to it. Depending on the printer, it may or may not print
	// to stdout.
	Write(in []byte)

	// Print prints the input to the printer formatting the received input
	// with `"FIRE" + join(<input>, " ") + "\n"`.
	Print(input ...string)
}

type DelegateToWriterPrinter struct {
	writer io.Writer
}

func (p *DelegateToWriterPrinter) Disabled() bool {
	return false
}

func (p *DelegateToWriterPrinter) Write(in []byte) {
	flushToFirehose(in, p.writer)
}

func (p *DelegateToWriterPrinter) Print(input ...string) {
	flushToFirehose([]byte("FIRE "+strings.Join(input, " ")+"\n"), p.writer)
}

// flushToFirehose sends data to Firehose via `io.Writter` checking for errors
// and retrying if necessary.
//
// If error is still present after 10 retries, prints an error message to `writer`
// as well as writing file `/tmp/firehose_writer_failed_print.log` with the same
// error message.
func flushToFirehose(in []byte, writer io.Writer) {
	var written int
	var err error
	loops := 10
	for i := 0; i < loops; i++ {
		written, err = writer.Write(in)

		if len(in) == written {
			return
		}

		in = in[written:]
		if i == loops-1 {
			break
		}
	}

	errstr := fmt.Sprintf("\nFIREHOSE FAILED WRITING %dx: %s\n", loops, err)
	os.WriteFile("/tmp/firehose_writer_failed_print.log", []byte(errstr), 0644)
	fmt.Fprint(writer, errstr)
}

type ToBufferPrinter struct {
	buffer *bytes.Buffer
}

func NewToBufferPrinter(initialAllocationSizeInBytes int) *ToBufferPrinter {
	return &ToBufferPrinter{
		buffer: bytes.NewBuffer(make([]byte, 0, initialAllocationSizeInBytes)),
	}
}

func NewToBufferPrinterWithBuffer(buffer *bytes.Buffer) *ToBufferPrinter {
	// Force a reset to ensure we start with a clean buffer
	buffer.Reset()

	return &ToBufferPrinter{
		buffer: buffer,
	}
}

func (p *ToBufferPrinter) Reset() {
	p.buffer.Reset()
}

func (p *ToBufferPrinter) Disabled() bool {
	return false
}

func (p *ToBufferPrinter) Write(in []byte) {
	p.buffer.Write(in)
}

func (p *ToBufferPrinter) Print(input ...string) {
	p.buffer.WriteString("FIRE " + strings.Join(input, " ") + "\n")
}

func (p *ToBufferPrinter) Buffer() *bytes.Buffer {
	return p.buffer
}

func Addr(in common.Address) string {
	return hex.EncodeToString(in[:])
}

func Bool(in bool) string {
	if in {
		return "true"
	}

	return "false"
}

func Hash(in common.Hash) string {
	return hex.EncodeToString(in[:])
}

func Hex(in []byte) string {
	if len(in) == 0 {
		return "."
	}

	return hex.EncodeToString(in)
}

func BigInt(in *big.Int) string {
	if in == nil {
		// This returns the same as if in would have been `big.NewInt(0)`
		return "."
	}

	return Hex(in.Bytes())
}

func Uint(in uint) string {
	return strconv.FormatUint(uint64(in), 10)
}

func Uint8(in uint8) string {
	return strconv.FormatUint(uint64(in), 10)
}

func Uint64(in uint64) string {
	return strconv.FormatUint(in, 10)
}

func JSON(in interface{}) string {
	out, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReportHeaderComparisonResult(actual *types.Header, expected *types.Header) {
	ReportToUser("There is a mismatch between Firehose genesis block and actual chain's stored genesis block, the actual genesis")
	ReportToUser("block's hash field extracted from Geth's database does not fit with hash of genesis block generated")
	ReportToUser("from Firehose determined genesis config, you might need to provide the correct 'genesis.json' file")
	ReportToUser("via --firehose-genesis-file")
	ReportToUser("")
	ReportToUser("Comparison of the actual Firehose recomputed genesis block <> expected Geth genesis block")

	compareAddress := fieldComparisonReporter(func(x interface{}) string { return x.(common.Address).String() })
	compareHash := fieldComparisonReporter(func(x interface{}) string { return x.(common.Hash).String() })
	compareUint64 := fieldComparisonReporter(func(x interface{}) string { return strconv.FormatUint(x.(uint64), 10) })
	compareBytes := fieldComparisonReporter(func(x interface{}) string { return hex.EncodeToString(x.([]byte)) })
	compareBigInt := fieldComparisonReporter(func(x interface{}) string {
		if x == nil || x.(*big.Int) == nil {
			return "<nil>"
		} else {
			return x.(*big.Int).String()
		}
	})

	compareHash("Hash", actual.Hash(), expected.Hash())
	compareUint64("Number", actual.Number.Uint64(), expected.Number.Uint64())
	compareHash("ParentHash", actual.ParentHash, expected.ParentHash)
	compareHash("UncleHash", actual.UncleHash, expected.UncleHash)
	compareAddress("Coinbase", actual.Coinbase, expected.Coinbase)
	compareHash("Root", actual.Root, expected.Root)
	compareHash("TxHash", actual.TxHash, expected.TxHash)
	compareHash("ReceiptHash", actual.ReceiptHash, expected.ReceiptHash)
	compareBytes("Bloom", actual.Bloom[:], expected.Bloom[:])
	compareBigInt("Difficulty", actual.Difficulty, expected.Difficulty)
	compareUint64("GasLimit", actual.GasLimit, expected.GasLimit)
	compareUint64("GasUsed", actual.GasUsed, expected.GasUsed)
	compareUint64("Time", actual.Time, expected.Time)
	compareBytes("Extra", actual.Extra, expected.Extra)
	compareHash("MixDigest", actual.MixDigest, expected.MixDigest)
	compareUint64("Nonce", actual.Nonce.Uint64(), expected.Nonce.Uint64())

	ReportToUser("")
}

func fieldComparisonReporter(toString func(x interface{}) string) func(field string, actual interface{}, expected interface{}) {
	return func(field string, actual interface{}, expected interface{}) {
		resolvedActual := toString(actual)
		resolvedExpected := toString(expected)

		sign := "!="
		if resolvedActual == resolvedExpected {
			sign = "=="
		}

		ReportToUser("%s [(actual) %s %s %s (expected)]", field, resolvedActual, sign, resolvedExpected)
	}
}

func ReportToUser(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
