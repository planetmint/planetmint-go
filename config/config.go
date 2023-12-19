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
stake-denom = "{{ .PlmntConfig.StakeDenom }}"
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
distribution-epochs = {{ .PlmntConfig.DistributionEpochs }}
re-issuance-epochs = {{ .PlmntConfig.ReIssuanceEpochs }}
`

// Config defines Planetmint's top level configuration
type Config struct {
	AssetRegistryScheme string `json:"asset-registry-scheme" mapstructure:"asset-registry-scheme"`
	AssetRegistryDomain string `json:"asset-registry-domain" mapstructure:"asset-registry-domain"`
	AssetRegistryPath   string `json:"asset-registry-path"   mapstructure:"asset-registry-path"`
	TokenDenom          string `json:"token-denom"           mapstructure:"token-denom"`
	StakeDenom          string `json:"stake-denom"           mapstructure:"stake-denom"`
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
	DistributionEpochs  int    `json:"distribution-epochs"   mapstructure:"distribution-epochs"`
	ReIssuanceEpochs    int    `json:"re-issuance-epochs"    mapstructure:"re-issuance-epochs"`
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
		StakeDenom:          "plmntstake",
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
		DistributionEpochs:  17640, // CometBFT epochs of 5s equate 1 day (12*60*24) + 15 min (15*24) to wait for confirmations on the re-issuance
		ReIssuanceEpochs:    17280, // CometBFT epochs of 5s equate 1 day (12*60*24)
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
