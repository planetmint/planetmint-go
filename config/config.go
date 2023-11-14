package config

import (
	"encoding/json"
	"sync"
)

const DefaultConfigTemplate = `
###############################################################################
###                         Planetmint                                      ###
###############################################################################

[planetmint]

asset-registry-endpoint = "{{ .PlmntConfig.AssetRegistryEndpoint }}"
token-denom = "{{ .PlmntConfig.TokenDenom }}"
stake-denom = "{{ .PlmntConfig.StakeDenom }}"
fee-denom = "{{ .PlmntConfig.FeeDenom }}"
pop-epochs = {{ .PlmntConfig.PoPEpochs }}
rpc-host = "{{ .PlmntConfig.RPCHost }}"
rpc-port = {{ .PlmntConfig.RPCPort }}
rpc-user = "{{ .PlmntConfig.RPCUser }}"
rpc-password = "{{ .PlmntConfig.RPCPassword }}"
mint-address = "{{ .PlmntConfig.MintAddress }}"
issuance-service-dir = "{{ .PlmntConfig.IssuanceServiceDir }}"
reissuance-asset = "{{ .PlmntConfig.ReissuanceAsset }}"
validator-address = "{{ .PlmntConfig.ValidatorAddress }}"
distribution-address-inv = "{{ .PlmntConfig.DistributionAddrInv }}"
distribution-address-dao = "{{ .PlmntConfig.DistributionAddrDAO }}"
distribution-address-pop = "{{ .PlmntConfig.DistributionAddrPoP }}"
distribution-epochs = {{ .PlmntConfig.DistributionEpochs }}

`

// Config defines Planetmint's top level configuration
type Config struct {
	AssetRegistryEndpoint string `mapstructure:"asset-registry-endpoint " json:"asset-registry-endpoint "`
	TokenDenom            string `mapstructure:"token-denom" json:"token-denom"`
	StakeDenom            string `mapstructure:"stake-denom" json:"stake-denom"`
	FeeDenom              string `mapstructure:"fee-denom" json:"fee-denom"`
	ConfigRootDir         string
	PoPEpochs             int    `mapstructure:"pop-epochs" json:"pop-epochs"`
	RPCHost               string `mapstructure:"rpc-host" json:"rpc-host"`
	RPCPort               int    `mapstructure:"rpc-port" json:"rpc-port"`
	RPCUser               string `mapstructure:"rpc-user" json:"rpc-user"`
	RPCPassword           string `mapstructure:"rpc-password" json:"rpc-password"`
	IssuanceServiceDir    string `mapstructure:"issuance-service-dir" json:"issuance-service-dir"`
	MintAddress           string `mapstructure:"mint-address" json:"mint-address"`
	ReissuanceAsset       string `mapstructure:"reissuance-asset" json:"reissuance-asset"`
	ValidatorAddress      string `mapstructure:"validator-address" json:"validator-address"`
	DistributionAddrInv   string `mapstructure:"distribution-address-inv" json:"distribution-address-inv"`
	DistributionAddrDAO   string `mapstructure:"distribution-address-dao" json:"distribution-address-dao"`
	DistributionAddrPoP   string `mapstructure:"distribution-address-pop" json:"distribution-address-pop"`
	DistributionEpochs    int    `mapstructure:"distribution-epochs" json:"distribution-epochs"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		AssetRegistryEndpoint: "https://assets.rddl.io/register_asset",
		TokenDenom:            "plmnt",
		StakeDenom:            "plmntstake",
		FeeDenom:              "plmnt",
		ConfigRootDir:         "",
		PoPEpochs:             24, // 24 CometBFT epochs of 5s equate 120s
		RPCHost:               "localhost",
		RPCPort:               18884,
		RPCUser:               "user",
		RPCPassword:           "password",
		IssuanceServiceDir:    "/opt/issuer_service",
		MintAddress:           "default",
		ReissuanceAsset:       "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
		ValidatorAddress:      "plmnt1w5dww335zhh98pzv783hqre355ck3u4w4hjxcx",
		DistributionAddrInv:   "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddrDAO:   "vjU8eMzU3JbUWZEpVANt2ePJuPWSPixgjiSj2jDMvkVVQQi2DDnZuBRVX4Ygt5YGBf5zvTWCr1ntdqYH",
		DistributionAddrPoP:   "vjTvXCFSReRsZ7grdsAreRR12KuKpDw8idueQJK9Yh1BYS7ggAqgvCxCgwh13KGK6M52y37HUmvr4GdD",
		DistributionEpochs:    17280, // CometBFT epochs of 5s equate 1 day (12*60*24)
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		plmntConfig = DefaultConfig()
	})
	return plmntConfig
}

func (config *Config) SetRoot(root string) *Config {
	config.ConfigRootDir = root
	return config
}

// SetWatchmenConfig sets Planetmint's configuration
func (config *Config) SetPlanetmintConfig(planetmintconfig interface{}) {
	jsonConfig, err := json.Marshal(planetmintconfig)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonConfig, &config)
	if err != nil {
		panic(err)
	}
}
