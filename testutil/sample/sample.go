package sample

import (
	"crypto/sha256"
	"encoding/hex"

	machinetypes "planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Mnemonic sample mnemonic to use in tests
const Mnemonic = "helmet hedgehog lab actor weekend elbow pelican valid obtain hungry rocket decade tower gallery fit practice cart cherry giggle hair snack glance bulb farm"

// PubKey corresponding public key to sample mnemonic
const PubKey = "02328de87896b9cbb5101c335f40029e4be898988b470abbf683f1a0b318d73470"

// Name is the name of the sample machine to use in tests
const Name = "machine"

// Amount is the amount to transfer to the machine account
const Amount = "1000stake"

// Fees is the amount of fees to use in tests
const Fees = "2stake"

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

func Machine(name, pubKey string) machinetypes.Machine {
	metadata := Metadata()
	m := machinetypes.Machine{
		Name:             name,
		Ticker:           name + "_ticker",
		Reissue:          true,
		Amount:           1000,
		Precision:        8,
		IssuerPlanetmint: pubKey,
		IssuerLiquid:     pubKey,
		MachineId:        pubKey,
		Metadata:         &metadata,
	}
	return m
}

func MachineIndex(pubKey string) machinetypes.MachineIndex {
	return machinetypes.MachineIndex{
		MachineId:        pubKey,
		IssuerPlanetmint: pubKey,
		IssuerLiquid:     pubKey,
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

func Asset(sk string) (string, string) {
	cid := "cid"

	skBytes, _ := hex.DecodeString(sk)
	privKey := &secp256k1.PrivKey{Key: skBytes}

	cidBytes, _ := hex.DecodeString(cid)
	hash := sha256.Sum256(cidBytes)

	sign, _ := privKey.Sign(hash[:])

	signatureHex := hex.EncodeToString(sign)

	return cid, signatureHex
}
