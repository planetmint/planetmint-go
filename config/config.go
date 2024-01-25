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
token-denom = "{{ .PlmntConfig.TokenDenom }}"
fee-denom = "{{ .PlmntConfig.FeeDenom }}"
staged-denom = "{{ .PlmntConfig.StagedDenom }}"
claim-denom = "{{ .PlmntConfig.ClaimDenom }}"
pop-epochs = {{ .PlmntConfig.PopEpochs }}
rpc-host = "{{ .PlmntConfig.RPCHost }}"
rpc-port = {{ .PlmntConfig.RPCPort }}
rpc-user = "{{ .PlmntConfig.RPCUser }}"
rpc-password = "{{ .PlmntConfig.RPCPassword }}"
rpc-scheme = "{{ .PlmntConfig.RPCScheme }}"
rpc-wallet = "{{ .PlmntConfig.RPCWallet }}"
reissuance-asset = "{{ .PlmntConfig.ReissuanceAsset }}"
validator-address = "{{ .PlmntConfig.ValidatorAddress }}"
distribution-address-inv = "{{ .PlmntConfig.DistributionAddrInv }}"
distribution-address-dao = "{{ .PlmntConfig.DistributionAddrDAO }}"
distribution-address-pop = "{{ .PlmntConfig.DistributionAddrPop }}"
distribution-offset = {{ .PlmntConfig.DistributionOffset }}
reissuance-epochs = {{ .PlmntConfig.ReissuanceEpochs }}
mqtt-domain = "{{ .PlmntConfig.MqttDomain }}"
mqtt-port = {{ .PlmntConfig.MqttPort }}
mqtt-user = "{{ .PlmntConfig.MqttUser }}"
mqtt-password = "{{ .PlmntConfig.MqttPassword }}"
mqtt-response-timeout = {{ .PlmntConfig.MqttResponseTimeout }}
`

// Config defines Planetmint's top level configuration
type Config struct {
	AssetRegistryScheme string `json:"asset-registry-scheme" mapstructure:"asset-registry-scheme"`
	AssetRegistryDomain string `json:"asset-registry-domain" mapstructure:"asset-registry-domain"`
	AssetRegistryPath   string `json:"asset-registry-path"   mapstructure:"asset-registry-path"`
	TokenDenom          string `json:"token-denom"           mapstructure:"token-denom"`
	FeeDenom            string `json:"fee-denom"             mapstructure:"fee-denom"`
	StagedDenom         string `json:"staged-denom"          mapstructure:"staged-denom"`
	ClaimDenom          string `json:"claim-denom"           mapstructure:"claim-denom"`
	ConfigRootDir       string `json:"config-root-dir"       mapstructure:"config-root-dir"`
	PopEpochs           int    `json:"pop-epochs"            mapstructure:"pop-epochs"`
	RPCHost             string `json:"rpc-host"              mapstructure:"rpc-host"`
	RPCPort             int    `json:"rpc-port"              mapstructure:"rpc-port"`
	RPCUser             string `json:"rpc-user"              mapstructure:"rpc-user"`
	RPCPassword         string `json:"rpc-password"          mapstructure:"rpc-password"`
	RPCScheme           string `json:"rpc-scheme"            mapstructure:"rpc-scheme"`
	RPCWallet           string `json:"rpc-wallet"            mapstructure:"rpc-wallet"`
	ReissuanceAsset     string `json:"reissuance-asset"      mapstructure:"reissuance-asset"`
	ValidatorAddress    string `json:"validator-address"     mapstructure:"validator-address"`
	DistributionAddrInv string `json:"distribution-addr-inv" mapstructure:"distribution-addr-inv"`
	DistributionAddrDAO string `json:"distribution-addr-dao" mapstructure:"distribution-addr-dao"`
	DistributionAddrPop string `json:"distribution-addr-pop" mapstructure:"distribution-addr-pop"`
	DistributionOffset  int    `json:"distribution-offset"   mapstructure:"distribution-offset"`
	ReissuanceEpochs    int    `json:"reissuance-epochs"     mapstructure:"reissuance-epochs"`
	MqttDomain          string `json:"mqtt-domain"           mapstructure:"mqtt-domain"`
	MqttPort            int    `json:"mqtt-port"             mapstructure:"mqtt-port"`
	MqttUser            string `json:"mqtt-user"             mapstructure:"mqtt-user"`
	MqttPassword        string `json:"mqtt-password"         mapstructure:"mqtt-password"`
	MqttResponseTimeout int    `json:"mqtt-response-timeout" mapstructure:"mqtt-response-timeout"`
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
		TokenDenom:          "plmnt",
		FeeDenom:            "plmnt",
		StagedDenom:         "stagedcrddl",
		ClaimDenom:          "crddl",
		ConfigRootDir:       "",
		PopEpochs:           24, // 24 CometBFT epochs of 5s equate 120s
		RPCHost:             "localhost",
		RPCPort:             18884,
		RPCUser:             "user",
		RPCPassword:         "password",
		RPCScheme:           "http",
		RPCWallet:           "rpcwallet",
		ReissuanceAsset:     "7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9",
		ValidatorAddress:    "plmnt1w5dww335zhh98pzv783hqre355ck3u4w4hjxcx",
		DistributionAddrInv: "vjTyRN2G42Yq3T5TJBecHj1dF1xdhKF89hKV4HJN3uXxUbaVGVR76hAfVRQqQCovWaEpar7G5qBBprFG",
		DistributionAddrDAO: "vjU8eMzU3JbUWZEpVANt2ePJuPWSPixgjiSj2jDMvkVVQQi2DDnZuBRVX4Ygt5YGBf5zvTWCr1ntdqYH",
		DistributionAddrPop: "vjTvXCFSReRsZ7grdsAreRR12KuKpDw8idueQJK9Yh1BYS7ggAqgvCxCgwh13KGK6M52y37HUmvr4GdD",
		// `DistributionOffset` relative to `ReissuanceEpochs`. CometBFT epochs of 5s equate 30 min (12*30)
		// to wait for confirmations on the reissuance
		DistributionOffset: 360,
		// `ReissuanceEpochs` is a configuration parameter that determines the number of CometBFT epochs
		// required for reissuance. In the context of Planetmint, reissuance refers to the process of
		// issuing new tokens. This configuration parameter specifies the number of epochs (each epoch is 5
		// seconds) that need to pass before reissuance can occur. In this case, `ReissuanceEpochs` is set
		// to 17280, which means that reissuance can occur after 1 day (12*60*24) of epochs.
		ReissuanceEpochs:    17280,
		MqttDomain:          "testnet-mqtt.rddl.io",
		MqttPort:            1885,
		MqttUser:            "user",
		MqttPassword:        "password",
		MqttResponseTimeout: 2000, // the value is defined in milliseconds
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
