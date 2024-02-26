package sample

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Mnemonic sample mnemonic to use in tests
const Mnemonic = "helmet hedgehog lab actor weekend elbow pelican valid obtain hungry rocket decade tower gallery fit practice cart cherry giggle hair snack glance bulb farm"

// PubKey corresponding public key to sample mnemonic
const PubKey = "021cd2a59c6f9402ce09effba89b3deb6bb5863733e625f22c06204918061db4f0"

// Name is the name of the sample machine to use in tests
const Name = "machine"

// FeeDenom is the fee denomination for e2e test cases
const FeeDenom = "plmnt"

// DefaultDerivationPath is the BIP44Prefix for PLMNT (see https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
const DefaultDerivationPath = "m/44'/8680'/0'/0/0"

// ConstBech32Addr constant bech32 address for mocks
const ConstBech32Addr = "plmnt10mq5nj8jhh27z7ejnz2ql3nh0qhzjnfvy50877"

// KeyPair returns a sample private / public keypair
func KeyPair(optional ...int) (string, string) {
	secret := "Don't tell anybody"
	// optional value if different keypairs are needed
	if len(optional) > 0 {
		secret = fmt.Sprintf("%v%v", secret, optional[0])
	}
	sk := secp256k1.GenPrivKeyFromSecret([]byte(secret))
	pk := sk.PubKey()
	return hex.EncodeToString(sk.Bytes()), hex.EncodeToString(pk.Bytes())
}

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func Secp256k1AccAddress() sdk.AccAddress {
	pk := secp256k1.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}

func Asset() string {
	cid := "cid0"
	return cid
}
