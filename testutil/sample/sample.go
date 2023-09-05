package sample

import (
	"encoding/hex"

	"planetmint-go/config"
	machinetypes "planetmint-go/x/machine/types"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
)

// Mnemonic sample mnemonic to use in tests
const Mnemonic = "helmet hedgehog lab actor weekend elbow pelican valid obtain hungry rocket decade tower gallery fit practice cart cherry giggle hair snack glance bulb farm"

// PubKey corresponding public key to sample mnemonic
const PubKey = "021cd2a59c6f9402ce09effba89b3deb6bb5863733e625f22c06204918061db4f0"

// Name is the name of the sample machine to use in tests
const Name = "machine"

// Amount is the amount to transfer to the machine account
const Amount = "1000stake"

// Fees is the amount of fees to use in tests
const Fees = "2stake"

// DefaultDerivationPath is the BIP44Prefix for PLMNT (see https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
const DefaultDerivationPath = "m/44'/8680'/0'/0/0"

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
	_, liquidPubKey := ExtendedKeyPair(config.LiquidNetParams)
	_, planetmintPubKey := ExtendedKeyPair(config.PlmntNetParams)
	m := machinetypes.Machine{
		Name:             name,
		Ticker:           name + "_ticker",
		Domain:           "lab.r3c.network",
		Reissue:          true,
		Amount:           1000,
		Precision:        8,
		IssuerPlanetmint: planetmintPubKey,
		IssuerLiquid:     liquidPubKey,
		MachineId:        pubKey,
		Metadata:         &metadata,
		Type:             1,
	}
	return m
}

func MachineIndex(pubKey string, planetmintPubKey string, liquidPubKey string) machinetypes.MachineIndex {
	return machinetypes.MachineIndex{
		MachineId:        pubKey,
		IssuerPlanetmint: planetmintPubKey,
		IssuerLiquid:     liquidPubKey,
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

	cid_bytes := []byte(cid)
	sign, _ := privKey.Sign(cid_bytes)

	signatureHex := hex.EncodeToString(sign)

	return cid, signatureHex
}

func ExtendedKeyPair(cfg chaincfg.Params) (string, string) {
	seed, err := bip39.NewSeedWithErrorChecking(Mnemonic, keyring.DefaultBIP39Passphrase)
	if err != nil {
		panic(err)
	}
	xprivKey, err := hdkeychain.NewMaster(seed, &cfg)
	if err != nil {
		panic(err)
	}
	xpubKey, err := xprivKey.Neuter()
	if err != nil {
		panic(err)
	}
	return xprivKey.String(), xpubKey.String()
}

func TrustAnchor() machinetypes.TrustAnchor {
	return machinetypes.TrustAnchor{
		Pubkey: PubKey,
	}
}
