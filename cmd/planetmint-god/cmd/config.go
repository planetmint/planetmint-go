package cmd

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/app"
)

const (
	// Address prefix suffixes
	pubKeySuffix     = "pub"
	valOperSuffix    = "valoper"
	valOperPubSuffix = "valoperpub"
	valConsSuffix    = "valcons"
	valConsPubSuffix = "valconspub"

	// PLMNT coin type as defined in SLIP44
	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	plmntCoinType = 8680
)

// initSDKConfig initializes and returns the SDK configuration with proper Bech32 prefixes
// and coin type settings for the Planetmint network.
func initSDKConfig() *sdk.Config {
	config := sdk.GetConfig()

	// Configure address prefixes
	configureBech32Prefixes(config)

	// Set coin type for PLMNT
	config.SetCoinType(plmntCoinType)

	// Seal the configuration to prevent further modifications
	config.Seal()

	return config
}

// configureBech32Prefixes sets up all the Bech32 prefixes for different address types
// using the base account address prefix defined in the app package.
func configureBech32Prefixes(config *sdk.Config) {
	// Account addresses
	config.SetBech32PrefixForAccount(
		app.AccountAddressPrefix,
		app.AccountAddressPrefix+pubKeySuffix,
	)

	// Validator addresses
	config.SetBech32PrefixForValidator(
		app.AccountAddressPrefix+valOperSuffix,
		app.AccountAddressPrefix+valOperPubSuffix,
	)

	// Consensus node addresses
	config.SetBech32PrefixForConsensusNode(
		app.AccountAddressPrefix+valConsSuffix,
		app.AccountAddressPrefix+valConsPubSuffix,
	)
}
