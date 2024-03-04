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
asset-registry-domain = "{{ .PlmntConfig.AssetRegistryDomain }}"
asset-registry-path = "{{ .PlmntConfig.AssetRegistryPath }}"
asset-registry-scheme = "{{ .PlmntConfig.AssetRegistryScheme}}"
claim-address = "{{ .PlmntConfig.ClaimAddress }}"
claim-denom = "{{ .PlmntConfig.ClaimDenom }}"
distribution-address-dao = "{{ .PlmntConfig.DistributionAddrDAO }}"
distribution-address-early-inv = "{{ .PlmntConfig.DistributionAddrEarlyInv }}"
distribution-address-investor = "{{ .PlmntConfig.DistributionAddrInvestor }}"
distribution-address-pop = "{{ .PlmntConfig.DistributionAddrPop }}"
distribution-address-strategic = "{{ .PlmntConfig.DistributionAddrStrategic }}"
distribution-offset = {{ .PlmntConfig.DistributionOffset }}
mqtt-response-timeout = {{ .PlmntConfig.MqttResponseTimeout }}
pop-epochs = {{ .PlmntConfig.PopEpochs }}
reissuance-asset = "{{ .PlmntConfig.ReissuanceAsset }}"
reissuance-epochs = {{ .PlmntConfig.ReissuanceEpochs }}
rpc-host = "{{ .PlmntConfig.RPCHost }}"
rpc-port = {{ .PlmntConfig.RPCPort }}
rpc-user = "{{ .PlmntConfig.RPCUser }}"
rpc-password = "{{ .PlmntConfig.RPCPassword }}"
rpc-scheme = "{{ .PlmntConfig.RPCScheme }}"
rpc-wallet = "{{ .PlmntConfig.RPCWallet }}"
staged-denom = "{{ .PlmntConfig.StagedDenom }}"
token-denom = "{{ .PlmntConfig.TokenDenom }}"
tx-gas-limit = {{ .PlmntConfig.TxGasLimit }}
validator-address = "{{ .PlmntConfig.ValidatorAddress }}"
mqtt-domain = "{{ .PlmntConfig.MqttDomain }}"
mqtt-port = {{ .PlmntConfig.MqttPort }}
mqtt-user = "{{ .PlmntConfig.MqttUser }}"
mqtt-password = "{{ .PlmntConfig.MqttPassword }}"
`

// Config defines Planetmint's top level configuration
type Config struct {
	AssetRegistryDomain       string `json:"asset-registry-domain"       mapstructure:"asset-registry-domain"`
	AssetRegistryPath         string `json:"asset-registry-path"         mapstructure:"asset-registry-path"`
	AssetRegistryScheme       string `json:"asset-registry-scheme"       mapstructure:"asset-registry-scheme"`
	ClaimAddress              string `json:"claim-address"               mapstructure:"claim-address"`
	ClaimDenom                string `json:"claim-denom"                 mapstructure:"claim-denom"`
	ConfigRootDir             string `json:"config-root-dir"             mapstructure:"config-root-dir"`
	DistributionAddrDAO       string `json:"distribution-addr-dao"       mapstructure:"distribution-addr-dao"`
	DistributionAddrEarlyInv  string `json:"distribution-addr-early-inv" mapstructure:"distribution-addr-early-inv"`
	DistributionAddrInvestor  string `json:"distribution-addr-investor"  mapstructure:"distribution-addr-investor"`
	DistributionAddrPop       string `json:"distribution-addr-pop"       mapstructure:"distribution-addr-pop"`
	DistributionAddrStrategic string `json:"distribution-addr-strategic" mapstructure:"distribution-addr-strategic"`
	DistributionOffset        int    `json:"distribution-offset"         mapstructure:"distribution-offset"`
	MqttDomain                string `json:"mqtt-domain"                 mapstructure:"mqtt-domain"`
	MqttPassword              string `json:"mqtt-password"               mapstructure:"mqtt-password"`
	MqttPort                  int    `json:"mqtt-port"                   mapstructure:"mqtt-port"`
	MqttResponseTimeout       int    `json:"mqtt-response-timeout"       mapstructure:"mqtt-response-timeout"`
	MqttUser                  string `json:"mqtt-user"                   mapstructure:"mqtt-user"`
	PopEpochs                 int    `json:"pop-epochs"                  mapstructure:"pop-epochs"`
	RPCHost                   string `json:"rpc-host"                    mapstructure:"rpc-host"`
	RPCPassword               string `json:"rpc-password"                mapstructure:"rpc-password"`
	RPCPort                   int    `json:"rpc-port"                    mapstructure:"rpc-port"`
	RPCScheme                 string `json:"rpc-scheme"                  mapstructure:"rpc-scheme"`
	RPCUser                   string `json:"rpc-user"                    mapstructure:"rpc-user"`
	RPCWallet                 string `json:"rpc-wallet"                  mapstructure:"rpc-wallet"`
	ReissuanceAsset           string `json:"reissuance-asset"            mapstructure:"reissuance-asset"`
	ReissuanceEpochs          int    `json:"reissuance-epochs"           mapstructure:"reissuance-epochs"`
	StagedDenom               string `json:"staged-denom"                mapstructure:"staged-denom"`
	TokenDenom                string `json:"token-denom"                 mapstructure:"token-denom"`
	TxGasLimit                int    `json:"tx-gas-limit"                mapstructure:"tx-gas-limit"`
	ValidatorAddress          string `json:"validator-address"           mapstructure:"validator-address"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		AssetRegistryDomain:       "testnet-assets.rddl.io",
		AssetRegistryPath:         "register_asset",
		AssetRegistryScheme:       "https",
		ClaimAddress:              "plmnt1w5dww335zhh98pzv783hqre355ck3u4w4hjxcx",
		ClaimDenom:                "crddl",
		ConfigRootDir:             "",
		DistributionAddrDAO:       "vjU8eMzU3JbUWZEpVANt2ePJuPWSPixgjiSj2jDMvkVVQQi2DDnZuBRVX4Ygt5YGBf5zvTWCr1ntdqYH",
		DistributionAddrEarlyInv:  "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddrInvestor:  "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddrPop:       "vjTvXCFSReRsZ7grdsAreRR12KuKpDw8idueQJK9Yh1BYS7ggAqgvCxCgwh13KGK6M52y37HUmvr4GdD",
		DistributionAddrStrategic: "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		// `DistributionOffset` relative to `ReissuanceEpochs`. CometBFT epochs of 5s equate 30 min (12*30)
		// to wait for confirmations on the reissuance
		DistributionOffset:  360,
		MqttDomain:          "testnet-mqtt.rddl.io",
		MqttPassword:        "password",
		MqttPort:            1885,
		MqttResponseTimeout: 2000, // the value is defined in milliseconds
		MqttUser:            "user",
		PopEpochs:           24, // 24 CometBFT epochs of 5s equate 120s
		RPCHost:             "localhost",
		RPCPassword:         "password",
		RPCPort:             18884,
		RPCScheme:           "http",
		RPCUser:             "user",
		RPCWallet:           "rpcwallet",
		ReissuanceAsset:     "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
		// `ReissuanceEpochs` is a configuration parameter that determines the number of CometBFT epochs
		// required for reissuance. In the context of Planetmint, reissuance refers to the process of
		// issuing new tokens. This configuration parameter specifies the number of epochs (each epoch is 5
		// seconds) that need to pass before reissuance can occur. In this case, `ReissuanceEpochs` is set
		// to 17280, which means that reissuance can occur after 1 day (12*60*24) of epochs.
		ReissuanceEpochs: 17280,
		StagedDenom:      "stagedcrddl",
		TokenDenom:       "plmnt",
		TxGasLimit:       200000,
		ValidatorAddress: "plmnt1w5dww335zhh98pzv783hqre355ck3u4w4hjxcx",
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

// CHANGE AGAIN
// func (config *Config) SetRoot(root string) *Config {
// 	config.ConfigRootDir = root
// 	return config
// }

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
