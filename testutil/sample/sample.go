package sample

import (
	"encoding/hex"
	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// KeyPair returns a sample private / public keypair
func KeyPair() (string, string) {
	secret := "Don't tell anybody"
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

func Machine(pkPM string, pkL string) machinetypes.Machine {
	m := machinetypes.Machine{
		Name:             "machine",
		Ticker:           "PM",
		Issued:           1,
		Precision:        8,
		IssuerPlanetmint: pkPM,
		IssuerLiquid:     pkL,
		Cid:              "Cid",
	}
	return m
}
