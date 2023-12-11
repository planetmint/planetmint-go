package sample

import (
	"encoding/hex"
	"fmt"

	"github.com/planetmint/planetmint-go/config"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

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
const Fees = "1stake"

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

// Machine creates a new machine object
// TODO: make address deterministic for test cases
func Machine(name, pubKey string, prvKey string, address string) machinetypes.Machine {
	metadata := Metadata()
	_, liquidPubKey := ExtendedKeyPair(config.LiquidNetParams)
	_, planetmintPubKey := ExtendedKeyPair(config.PlmntNetParams)

	prvKeyBytes, _ := hex.DecodeString(prvKey)
	sk := &secp256k1.PrivKey{Key: prvKeyBytes}
	pubKeyBytes, _ := hex.DecodeString(pubKey)
	sign, _ := sk.Sign(pubKeyBytes)
	signatureHex := hex.EncodeToString(sign)

	if address == "" {
		address = Secp256k1AccAddress().String()
	}

	m := machinetypes.Machine{
		Name:               name,
		Ticker:             name + "_ticker",
		Domain:             "lab.r3c.network",
		Reissue:            true,
		Amount:             1000,
		Precision:          8,
		IssuerPlanetmint:   planetmintPubKey,
		IssuerLiquid:       liquidPubKey,
		MachineId:          pubKey,
		Metadata:           &metadata,
		Type:               1,
		MachineIdSignature: signatureHex,
		Address:            address,
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

func Asset() string {
	cid := "cid0"
	return cid
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

func TrustAnchor(pubkey string) machinetypes.TrustAnchor {
	return machinetypes.TrustAnchor{
		Pubkey: pubkey,
	}
}

func MintRequest(beneficiaryAddr string, amount uint64, txhash string) daotypes.MintRequest {
	return daotypes.MintRequest{
		Beneficiary:  beneficiaryAddr,
		Amount:       amount,
		LiquidTxHash: txhash,
	}
}
