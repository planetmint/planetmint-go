package config

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// LiquidNetParams defines the network parameters for the Liquid network.
var LiquidNetParams chaincfg.Params

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
	// Not allowed to register LiquidNetParams, because it's just another
	// Bitcoin network with different coin type.
	// See https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	LiquidNetParams = chaincfg.MainNetParams
	LiquidNetParams.Name = "liquidv1"
	LiquidNetParams.HDCoinType = 1776

	// Need to register PlmntNetParams, otherwise we get an "unknown hd
	// private extended key bytes" error.
	err := chaincfg.Register(&PlmntNetParams)
	if err != nil {
		panic(err)
	}
}
