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
osc-service-port = {{ .PlmntConfig.OSCServicePort }}
watchmen-endpoint = "{{ .PlmntConfig.WatchmenEndpoint }}"
watchmen-port = {{ .PlmntConfig.WatchmenPort }}
token-denom = "{{ .PlmntConfig.TokenDenom }}"
stake-denom = "{{ .PlmntConfig.StakeDenom }}"
fee-denom = "{{ .PlmntConfig.FeeDenom }}"
config-root-dir = "{{ .PlmntConfig.ConfigRootDir }}"
pop-epochs = {{ .PlmntConfig.PoPEpochs }}
rpc-host = "{{ .PlmntConfig.RPCHost }}"
rpc-port = {{ .PlmntConfig.RPCPort }}
rpc-user = "{{ .PlmntConfig.RPCUser }}"
rpc-password = "{{ .PlmntConfig.RPCPassword }}"
`

// Config defines Planetmint's top level configuration
type Config struct {
	OSCServicePort   int    `mapstructure:"osc-service-port" json:"osc-service-port"`
	WatchmenEndpoint string `mapstructure:"watchmen-endpoint" json:"watchmen-endpoint"`
	WatchmenPort     int    `mapstructure:"watchmen-port" json:"watchmen-port"`
	TokenDenom       string `mapstructure:"token-denom" json:"token-denom"`
	StakeDenom       string `mapstructure:"stake-denom" json:"stake-denom"`
	FeeDenom         string `mapstructure:"fee-denom" json:"fee-denom"`
	ConfigRootDir    string `mapstructure:"config-root-dir" json:"config-root-dir"`
	PoPEpochs        int    `mapstructure:"pop-epochs" json:"pop-epochs"`
	RPCHost          string `mapstructure:"rpc-host" json:"rpc-host"`
	RPCPort          int    `mapstructure:"rpc-port" json:"rpc-port"`
	RPCUser          string `mapstructure:"rpc-user" json:"rpc-user"`
	RPCPassword      string `mapstructure:"rpc-password" json:"rpc-password"`
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
		OSCServicePort:   8766,
		WatchmenEndpoint: "lab.r3c.network",
		WatchmenPort:     7401,
		TokenDenom:       "plmnt",
		StakeDenom:       "plmntstake",
		FeeDenom:         "plmnt",
		ConfigRootDir:    filepath.Join(currentUser.HomeDir, ".planetmint-go"),
		PoPEpochs:        24, // 24 CometBFT epochs of 5s equate 120s
		RPCHost:          "localhost",
		RPCPort:          18884,
		RPCUser:          "user",
		RPCPassword:      "password",
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
