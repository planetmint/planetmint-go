package sample

import (
	"encoding/hex"
	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// KeyPair returns a sample private / public keypair
func KeyPair() (string, string) {
	secret := "Don't tell anybody"
	sk := ed25519.GenPrivKeyFromSecret([]byte(secret))
	pk := sk.PubKey()
	return hex.EncodeToString(sk.Bytes()), hex.EncodeToString(pk.Bytes())
}

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

func Machine(machineId string, pkPM string, pkL string) machinetypes.Machine {
	m := machinetypes.Machine{
		Name:             "machine",
		Ticker:           "PM",
		Issued:           1,
		Amount:           1000,
		Precision:        8,
		IssuerPlanetmint: pkPM,
		IssuerLiquid:     pkL,
		MachineId:        machineId,
	}
	return m
}

func MachineIndex(machineId string, pkPM string, pkL string) machinetypes.MachineIndex {
	return machinetypes.MachineIndex{
		MachineId:        machineId,
		IssuerPlanetmint: pkPM,
		IssuerLiquid:     pkL,
	}
}

func Metadata() machinetypes.Metadata {
	return machinetypes.Metadata{
		Gps:               "{\"Latitude\":\"-48.876667\",\"Longitude\":\"-123.393333\"}",
		Device:            "{\"Manufacturer\": \"RDDL\",\"Serial\":\"AdnT2uyt\"}",
		AssetDefinition:   "{\"Version\": \"0.1\"}",
		AdditionalDataCID: "CID",
	}
}
