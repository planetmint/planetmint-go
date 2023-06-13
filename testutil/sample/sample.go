package sample

import (
	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// // KeyPair returns a sample private / public keypair
// func KeyPair() (ed25519.PrivKey, cryptotypes.PubKey) {
// 	sk := ed25519.GenPrivKey()
// 	pk := sk.PubKey()
// 	return sk.Key, pk
// }

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func Machine() machinetypes.Machine {
	m := machinetypes.Machine{
		Name:             "machine",
		Ticker:           "PM",
		Issued:           1,
		Precision:        8,
		IssuerPlanetmint: "pubkey",
		IssuerLiquid:     "pubkey",
		Cid:              "Cid",
	}
	return m
}
