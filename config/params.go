package config

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// PlmntNetParams defines the network parameters for the Planetmint network.
var PlmntNetParams = chaincfg.Params{
	Name: "planetmint",

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x03, 0xe1, 0x42, 0xb0}, // starts with pmpr
	HDPublicKeyID:  [4]byte{0x03, 0xe1, 0x42, 0x47}, // starts with pmpb

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 8680,
}

func init() {
	chaincfg.Register(&PlmntNetParams)
}
