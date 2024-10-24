package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
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
	return NewParams("https", "default-param-set.rddl.io", "default_param_set_register_asset", 8800, "plmnt")
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
