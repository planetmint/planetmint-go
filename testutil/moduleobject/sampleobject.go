package moduleobject

import (
	"encoding/hex"
	"strconv"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/planetmint/planetmint-go/config"
	"github.com/planetmint/planetmint-go/testutil/sample"
	daotypes "github.com/planetmint/planetmint-go/x/dao/types"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/go-bip39"
)

func ExtendedKeyPair(cfg chaincfg.Params) (string, string) {
	seed, err := bip39.NewSeedWithErrorChecking(sample.Mnemonic, keyring.DefaultBIP39Passphrase)
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
		address = sample.Secp256k1AccAddress().String()
	}

	m := machinetypes.Machine{
		Name:               name,
		Ticker:             name + "_ticker",
		Domain:             "testnet-assets.rddl.iok",
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

func MachineRandom(name, pubKey string, prvKey string, address string, random int) machinetypes.Machine {
	metadata := Metadata()
	_, liquidPubKey := ExtendedKeyPair(config.LiquidNetParams)
	_, planetmintPubKey := ExtendedKeyPair(config.PlmntNetParams)

	prvKeyBytes, _ := hex.DecodeString(prvKey)
	sk := &secp256k1.PrivKey{Key: prvKeyBytes}
	pubKeyBytes, _ := hex.DecodeString(pubKey)
	sign, _ := sk.Sign(pubKeyBytes)
	signatureHex := hex.EncodeToString(sign)

	if address == "" {
		address = sample.Secp256k1AccAddress().String()
	}

	m := machinetypes.Machine{
		Name:               name + strconv.Itoa(random),
		Ticker:             name + strconv.Itoa(random) + "_ticker",
		Domain:             "testnet-assets.rddl.io",
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
