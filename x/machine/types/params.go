package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(assetRegistryScheme string, assetRegistryDomain string, assetRegistryPath string,
	daoMachineFundingAmount uint64, daoMachineFundingDenom string) Params {
	return Params{
		AssetRegistryScheme:     assetRegistryScheme,
		AssetRegistryDomain:     assetRegistryDomain,
		AssetRegistryPath:       assetRegistryPath,
		DaoMachineFundingAmount: daoMachineFundingAmount,
		DaoMachineFundingDenom:  daoMachineFundingDenom,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("https", "testnet-assets.rddl.io", "register_asset", 8800, "plmnt")
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
