package config

import (
	"encoding/json"
	"fmt"
	"sync"
)

const DefaultConfigTemplate = `
###############################################################################
###                         Planetmint                                      ###
###############################################################################

[planetmint]

asset-registry-scheme = "{{ .PlmntConfig.AssetRegistryScheme}}"
asset-registry-domain = "{{ .PlmntConfig.AssetRegistryDomain }}"
asset-registry-path = "{{ .PlmntConfig.AssetRegistryPath }}"
fee-denom = "{{ .PlmntConfig.FeeDenom }}"
rpc-host = "{{ .PlmntConfig.RPCHost }}"
rpc-port = {{ .PlmntConfig.RPCPort }}
rpc-user = "{{ .PlmntConfig.RPCUser }}"
rpc-password = "{{ .PlmntConfig.RPCPassword }}"
rpc-scheme = "{{ .PlmntConfig.RPCScheme }}"
rpc-wallet = "{{ .PlmntConfig.RPCWallet }}"
validator-address = "{{ .PlmntConfig.ValidatorAddress }}"
mqtt-domain = "{{ .PlmntConfig.MqttDomain }}"
mqtt-port = {{ .PlmntConfig.MqttPort }}
mqtt-user = "{{ .PlmntConfig.MqttUser }}"
mqtt-password = "{{ .PlmntConfig.MqttPassword }}"

`

// Config defines Planetmint's top level configuration
type Config struct {
	AssetRegistryScheme string `json:"asset-registry-scheme" mapstructure:"asset-registry-scheme"`
	AssetRegistryDomain string `json:"asset-registry-domain" mapstructure:"asset-registry-domain"`
	AssetRegistryPath   string `json:"asset-registry-path"   mapstructure:"asset-registry-path"`
	FeeDenom            string `json:"fee-denom"             mapstructure:"fee-denom"`
	ConfigRootDir       string `json:"config-root-dir"       mapstructure:"config-root-dir"`
	RPCHost             string `json:"rpc-host"              mapstructure:"rpc-host"`
	RPCPort             int    `json:"rpc-port"              mapstructure:"rpc-port"`
	RPCUser             string `json:"rpc-user"              mapstructure:"rpc-user"`
	RPCPassword         string `json:"rpc-password"          mapstructure:"rpc-password"`
	RPCScheme           string `json:"rpc-scheme"            mapstructure:"rpc-scheme"`
	RPCWallet           string `json:"rpc-wallet"            mapstructure:"rpc-wallet"`
	ValidatorAddress    string `json:"validator-address"     mapstructure:"validator-address"`
	MqttDomain          string `json:"mqtt-domain"           mapstructure:"mqtt-domain"`
	MqttPort            int    `json:"mqtt-port"             mapstructure:"mqtt-port"`
	MqttUser            string `json:"mqtt-user"             mapstructure:"mqtt-user"`
	MqttPassword        string `json:"mqtt-password"         mapstructure:"mqtt-password"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		AssetRegistryScheme: "https",
		AssetRegistryDomain: "testnet-assets.rddl.io",
		AssetRegistryPath:   "register_asset",
		FeeDenom:            "plmnt",
		ConfigRootDir:       "",
		RPCHost:             "localhost",
		RPCPort:             18884,
		RPCUser:             "user",
		RPCPassword:         "password",
		RPCScheme:           "http",
		RPCWallet:           "rpcwallet",
		ValidatorAddress:    "plmnt1w5dww335zhh98pzv783hqre355ck3u4w4hjxcx",
		MqttDomain:          "testnet-mqtt.rddl.io",
		MqttPort:            1885,
		MqttUser:            "user",
		MqttPassword:        "password",
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		plmntConfig = DefaultConfig()
	})
	return plmntConfig
}

// GetRPCURL returns the elements RPC URL
func (config *Config) GetRPCURL() (url string) {
	url = fmt.Sprintf("%s://%s:%s@%s:%d/wallet/%s", config.RPCScheme, config.RPCUser, config.RPCPassword, config.RPCHost, config.RPCPort, config.RPCWallet)
	return
}

func (config *Config) SetRoot(root string) *Config {
	config.ConfigRootDir = root
	return config
}

// SetPlanetmintConfig sets Planetmint's configuration
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
