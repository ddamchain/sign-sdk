package account

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

const defaultPath = "m/44'/378'/0'/0/%d" // 372

type Master struct {
	key *hdkeychain.ExtendedKey
}

func NewMaster(mnemonic string) *Master {
	if mnemonic == "" {
		return nil
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil
	}
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil
	}
	return &Master{masterKey}
}

func (m Master) GetChild(index int) (*Child, error) {
	path, err := parseDerivationPath(fmt.Sprintf(defaultPath, index))
	if err != nil {
		return nil, err
	}
	key := m.key
	for _, n := range path {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}
	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}
	privateKeyECDSA := privateKey.ToECDSA()
	return NewChild(privateKeyECDSA.D.Bytes()), nil
}

func (m Master) TryGetChild(index int) *Child {
	child, err := m.GetChild(index)
	if err != nil {
		return nil
	}
	return child
}
