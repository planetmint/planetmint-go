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
func NewParams(mintAddress string, tokenDenom string, stagedDenom string,
	claimDenom string, reissuanceAsset string, reissuanceEpochs int64, popEpochs int64,
	distributionOffset int64, distributionAddressEarlyInv string, distributionAddressInvestor string,
	distributionAddressStrategic string, distributionAddressDao string, distributionAddressPop string,
	mqttResponseTimeout int64, claimAddress string, txGasLimit uint64, validatorPoPReward uint64) Params {
	return Params{
		MintAddress:     mintAddress,
		TokenDenom:      tokenDenom,
		StagedDenom:     stagedDenom,
		ClaimDenom:      claimDenom,
		ReissuanceAsset: reissuanceAsset,
		// `ReissuanceEpochs` is a configuration parameter that determines the number of CometBFT epochs
		// required for reissuance. In the context of Planetmint, reissuance refers to the process of
		// issuing new tokens. This configuration parameter specifies the number of epochs (each epoch is 5
		// seconds) that need to pass before reissuance can occur. In this case, `ReissuanceEpochs` is set
		// to 17280, which means that reissuance can occur after 1 day (12*60*24) of epochs.
		ReissuanceEpochs:   reissuanceEpochs,
		PopEpochs:          popEpochs,
		DistributionOffset: distributionOffset,
		// `DistributionOffset` relative to `ReissuanceEpochs`. CometBFT epochs of 5s equate 30 min (12*30)
		// to wait for confirmations on the reissuance
		DistributionAddressEarlyInv:  distributionAddressEarlyInv,
		DistributionAddressInvestor:  distributionAddressInvestor,
		DistributionAddressStrategic: distributionAddressStrategic,
		DistributionAddressDao:       distributionAddressDao,
		DistributionAddressPop:       distributionAddressPop,
		MqttResponseTimeout:          mqttResponseTimeout,
		ClaimAddress:                 claimAddress,
		TxGasLimit:                   txGasLimit,
		ValidatorPopReward:           validatorPoPReward,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		"plmnt1dyuhg8ldu3d6nvhrvzzemtc3893dys9v9lvdty",
		"plmnt",
		"stagedcrddl",
		"crddl",
		"not_a_valid_asset_id_default_param_set",
		3600,
		5,
		75,
		"early_inv_address_default_param_set",
		"investor_address_default_param_set",
		"strategic_address_default_param_set",
		"dao_address_default_param_set",
		"pop_address_default_param_set",
		2000,
		"plmnt1m5apfematgm7uueazhk482026ert95x2l2dx78",
		200000,
		100000000,
	)
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
