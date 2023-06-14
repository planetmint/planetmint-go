package sample

import (
	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// KeyPair returns a sample private / public keypair
func KeyPair() (string, string) {
	secret := "Hello World!"
	sk := ed25519.GenPrivKeyFromSecret([]byte(secret))
	pk := sk.PubKey()

	return sk.String(), pk.String()
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
