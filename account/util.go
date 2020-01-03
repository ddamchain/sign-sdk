package account

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sign-sdk/common"
	"sign-sdk/types"
	"strings"
)

const (
	HexPrefix     = "0x"
	AddrPrefix    = "DD"
	AddressLength = 32 //Length of Address( golang.SHA3，256-bit)
	HashLength    = 32 //Length of Hash (golang.SHA3, 256-bit)。
)

var DefaultRootDerivationPath = []uint32{0x80000000 + 44, 0x80000000 + 60, 0x80000000 + 0, 0}

func parseDerivationPath(path string) ([]uint32, error) {
	var result []uint32

	// Handle absolute or relative paths
	components := strings.Split(path, "/")
	switch {
	case len(components) == 0:
		return nil, errors.New("empty derivation path")

	case strings.TrimSpace(components[0]) == "":
		return nil, errors.New("ambiguous path: use 'm/' prefix for absolute paths, or no leading '/' for relative ones")

	case strings.TrimSpace(components[0]) == "m":
		components = components[1:]

	default:
		result = append(result, DefaultRootDerivationPath...)
	}
	// All remaining components are relative, append one by one
	if len(components) == 0 {
		return nil, errors.New("empty derivation path") // Empty relative paths
	}
	for _, component := range components {
		// Ignore any user added whitespace
		component = strings.TrimSpace(component)
		var value uint32

		// Handle hardened paths
		if strings.HasSuffix(component, "'") {
			value = 0x80000000
			component = strings.TrimSpace(strings.TrimSuffix(component, "'"))
		}
		// Handle the non hardened component
		bigval, ok := new(big.Int).SetString(component, 0)
		if !ok {
			return nil, fmt.Errorf("invalid component: %s", component)
		}
		max := math.MaxUint32 - value
		if bigval.Sign() < 0 || bigval.Cmp(big.NewInt(int64(max))) > 0 {
			if value == 0 {
				return nil, fmt.Errorf("component %v out of allowed range [0, %d]", bigval, max)
			}
			return nil, fmt.Errorf("component %v out of allowed hardened range [0, %d]", bigval, max)
		}
		value += uint32(bigval.Uint64())

		// Append and repeat
		result = append(result, value)
	}
	return result, nil
}

type bufferWriter struct {
	buf bytes.Buffer
}

func (bw *bufferWriter) writeByte(b byte) {
	bw.buf.WriteByte(b)
}

func (bw *bufferWriter) writeBytes(b []byte) {
	// Write len with big-endian
	bw.buf.Write(common.Int32ToByte(int32(len(b))))
	if len(b) > 0 {
		bw.buf.Write(b)
	}
}

func (bw *bufferWriter) Bytes() []byte {
	return bw.buf.Bytes()
}

func StringToAddress(s string) common.Address {
	if len(s) > len(AddrPrefix) {
		if AddrPrefix == strings.ToLower(s[0:len(AddrPrefix)]) {
			s = s[len(AddrPrefix):]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
	}
	bs, _ := hex.DecodeString(s)
	return BytesToAddress(bs)
}
func BytesToAddress(b []byte) common.Address {
	var a common.Address
	a.SetBytes(b)
	return a
}
func SetBytes(a *common.Address, b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[:], b[:])
}

func txRawToTransaction(targetAddr string, Value int64, GasLimit int64, GasPrice int64, Type int, Nonce int64, dataHex string) *Transaction {
	var target *common.Address
	if targetAddr != "" {
		t := StringToAddress(targetAddr)
		target = &t
	}

	data := make([]byte, 0)
	if dataHex != "" {
		//data,err:=base64.StdEncoding.DecodeString(dataHex)
		result, err := dataDecode(dataHex)
		if err != nil {
			fmt.Println("===========================> dataDecode err :", err)
			return nil
		}
		data = result
	}

	return &Transaction{
		Data:     data,
		Value:    types.NewBigInt(uint64(Value)),
		Nonce:    uint64(Nonce),
		Target:   target,
		Type:     int8(Type),
		GasLimit: types.NewBigInt(uint64(GasLimit)),
		GasPrice: types.NewBigInt(uint64(GasPrice)),
		//Sign:     crypto.HexToSign(tx.Sign),
	}
}

type Transaction struct {
	Data   []byte          `msgpack:"dt,omitempty"` // Data of the transaction, cost gas
	Value  *types.BigInt   `msgpack:"v"`            // The value the sender suppose to transfer
	Nonce  uint64          `msgpack:"nc"`           // The nonce indicates the transaction sequence related to sender
	Target *common.Address `msgpack:"tg,omitempty"` // The receiver address
	Type   int8            `msgpack:"tp"`           // Transaction type

	GasLimit *types.BigInt `msgpack:"gl"`
	GasPrice *types.BigInt `msgpack:"gp"`
	Hash     common.Hash   `msgpack:"h"`

	Sign   *common.Sign    `msgpack:"si"`  // The Sign of the sender
	Source *common.Address `msgpack:"src"` // Sender address, recovered from sign
}

func (tx *Transaction) GenHash() common.Hash {
	if nil == tx {
		return common.Hash{}
	}

	var (
		target   []byte
		value    []byte
		gasLimit []byte
		gasPrice []byte
	)
	if tx.Target != nil {
		target = tx.Target.Bytes()
	}
	if tx.Value != nil {
		value = tx.Value.Bytes()
	}
	if tx.GasLimit != nil {
		gasLimit = tx.GasLimit.Bytes()
	}
	if tx.GasPrice != nil {
		gasPrice = tx.GasPrice.Bytes()
	}

	txH := &txHashing{
		target:   target,
		value:    value,
		gasLimit: gasLimit,
		gasPrice: gasPrice,
		nonce:    new(big.Int).SetUint64(tx.Nonce).Bytes(),
		typ:      byte(tx.Type),
		data:     tx.Data,
	}

	return txH.genHash()
}

type txHashing struct {
	target   []byte
	value    []byte // bytes with big-endian
	gasLimit []byte // bytes with big-endian
	gasPrice []byte // bytes with big-endian
	nonce    []byte // bytes with big-endian
	typ      byte
	data     []byte
}

func (th *txHashing) genHash() common.Hash {
	buf := &bufferWriter{}
	buf.writeBytes(th.target)
	buf.writeBytes(th.value)
	buf.writeBytes(th.nonce)
	buf.writeBytes(th.gasLimit)
	buf.writeBytes(th.gasPrice)
	buf.writeByte(th.typ)
	buf.writeBytes(th.data)
	return common.BytesToHash(common.Sha256(buf.Bytes()))
}
