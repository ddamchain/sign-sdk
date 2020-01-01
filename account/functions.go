package account

import (
	"C"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tyler-smith/go-bip39"
)

func Sk2Pk(sk string) string {
	bs, err := hexutil.Decode(sk)
	if err != nil {
		return ""
	}
	child := NewChild(bs)
	if child == nil {
		return ""
	}
	return child.PublicKeyHex()
}

func Sk2Addr(sk string) string {
	bs, err := hexutil.Decode(sk)
	if err != nil {
		return ""
	}
	child := NewChild(bs)
	if child == nil {
		return ""
	}
	return child.AddressHex()
}

//export SignTx
func SignTx(sk string, targetAddr string, value int64, gasLimit int64, gasPrice int64, _type int, nonce int64, dataHex string) string {
	bs, err := hexutil.Decode(sk)
	if err != nil {
		return ""
	}
	child := NewChild(bs)
	if child == nil {
		return ""
	}
	return child.Sign(targetAddr, value, gasLimit, gasPrice, _type, nonce, dataHex)
}

func GenMnemonic() string {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return ""
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return ""
	}
	return mnemonic
}

func GetChild(mnemonic string, index int) string {
	if mnemonic == "" {
		return ""
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return ""
	}

	master := NewMaster(mnemonic)
	if master == nil {
		return ""
	}
	child, err := master.GetChild(index)
	if err != nil {
		return ""
	}
	return child.PrivateKeyHex()
}
