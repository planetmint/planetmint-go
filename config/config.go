package config

import (
	"encoding/json"
	"os/user"
	"path/filepath"
	"sync"
)

const DefaultConfigTemplate = `
###############################################################################
###                         Planetmint                                      ###
###############################################################################

[planetmint]

asset-registry-endpoint = "{{ .PlmntConfig.AssetRegistryEndpoint }}"
watchmen-port = {{ .PlmntConfig.WatchmenPort }}
token-denom = "{{ .PlmntConfig.TokenDenom }}"
stake-denom = "{{ .PlmntConfig.StakeDenom }}"
fee-denom = "{{ .PlmntConfig.FeeDenom }}"
config-root-dir = "{{ .PlmntConfig.ConfigRootDir }}"
pop-epochs = {{ .PlmntConfig.PoPEpochs }}
issuance-endpoint = "{{ .PlmntConfig.IssuanceEndpoint }}"
issuance-port = {{ .PlmntConfig.IssuancePort }}
`

// Config defines Planetmint's top level configuration
type Config struct {
	AssetRegistryEndpoint string `mapstructure:"asset-registry-endpoint " json:"asset-registry-endpoint "`
	WatchmenPort          int    `mapstructure:"watchmen-port" json:"watchmen-port"`
	TokenDenom            string `mapstructure:"token-denom" json:"token-denom"`
	StakeDenom            string `mapstructure:"stake-denom" json:"stake-denom"`
	FeeDenom              string `mapstructure:"fee-denom" json:"fee-denom"`
	ConfigRootDir         string `mapstructure:"config-root-dir" json:"config-root-dir"`
	PoPEpochs             int    `mapstructure:"pop-epochs" json:"pop-epochs"`
	IssuanceEndpoint      string `mapstructure:"issuance-endpoint" json:"issuance-endpoint"`
	IssuancePort          int    `mapstructure:"issuance-port" json:"issuance-port"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	return &Config{
		AssetRegistryEndpoint: "https://assets.rddl.io/register_asset",
		WatchmenPort:          7401,
		TokenDenom:            "plmnt",
		StakeDenom:            "plmntstake",
		FeeDenom:              "plmnt",
		ConfigRootDir:         filepath.Join(currentUser.HomeDir, ".planetmint-go"),
		PoPEpochs:             24, // 24 CometBFT epochs of 5s equate 120s
		IssuanceEndpoint:      "lab.r3c.network",
		IssuancePort:          7401,
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		plmntConfig = DefaultConfig()
	})
	return plmntConfig
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
