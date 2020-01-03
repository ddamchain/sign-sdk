package account

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
	"math/big"
	"sign-sdk/common/secp256k1"
)

const SignLength = 65 //length of signature，32 bytes r & 32 bytes s & 1 byte recid.

type Child struct {
	pk *ecdsa.PrivateKey
}

func (child Child) PrivateKeyHex() string {
	return hexutil.Encode(child.pk.D.Bytes())
}

func (child Child) PublicKeyHex() string {
	buf := elliptic.Marshal(child.pk.PublicKey.Curve, child.pk.PublicKey.X, child.pk.PublicKey.Y)
	return hexutil.Encode(buf)
}

func (child Child) AddressHex() string {
	x := child.pk.PublicKey.X.Bytes()
	y := child.pk.PublicKey.Y.Bytes()
	x = append(x, y...)

	addrBuf := sha3.Sum256(x)

	return encode2(addrBuf[:])
}

func encode2(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "DD")
	hex.Encode(enc[2:], b)
	return string(enc)
}

func (child Child) sign(hash []byte) *Sign {
	pribytes := child.pk.D.Bytes()
	seckbytes := pribytes
	if len(pribytes) < 32 {
		seckbytes = make([]byte, 32)
		copy(seckbytes[32-len(pribytes):32], pribytes) //make sure that the length of seckey is 32 bytes
	}

	sig, err := secp256k1.Sign(hash, seckbytes)
	if err == nil {
		return newSign(sig)
	} else {
		return nil
	}
}

func dataDecode(data string) ([]byte, error) {
	var a []byte
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func (child Child) Sign(targetAddr string, value int64, gasLimit int64, gasPrice int64, _type int, nonce int64, dataHex string) string {
	//trans:=txRawToTransaction(targetAddr,uint64(value),uint64(gasLimit),uint64(gasPrice),_type,uint64(nonce),dataHex)
	trans := txRawToTransaction(targetAddr, value, gasLimit, gasPrice, _type, nonce, dataHex)
	hash := trans.GenHash()

	sign := child.sign(hash.Bytes())

	return hexutil.Encode(sign.Bytes())
}

func NewChild(key []byte) *Child {
	child := Child{&ecdsa.PrivateKey{}}
	var one = new(big.Int).SetInt64(1)

	params := secp256k1.S256().Params()
	d := new(big.Int).SetBytes(key)
	if d.Cmp(params.N) >= 0 || d.Cmp(one) < 0 {
		return nil
	}

	child.pk.Curve = secp256k1.S256()
	child.pk.D = d
	child.pk.PublicKey.X, child.pk.PublicKey.Y = child.pk.Curve.ScalarBaseMult(key)
	return &child
}

// Sign Data struct
type Sign struct {
	r     big.Int
	s     big.Int
	recid byte
}

func newSign(b []byte) *Sign {
	if len(b) == 65 {
		var r, s big.Int
		br := b[:32]
		r = *r.SetBytes(br)

		sr := b[32:64]
		s = *s.SetBytes(sr)

		recid := b[64]
		return &Sign{r, s, recid}
	}
	return nil
}

func (s Sign) Bytes() []byte {
	rb := s.r.Bytes()
	sb := s.s.Bytes()
	r := make([]byte, SignLength)
	copy(r[32-len(rb):32], rb)
	copy(r[64-len(sb):64], sb)
	r[64] = s.recid
	return r
}
