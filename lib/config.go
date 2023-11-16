package lib

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/lib/params"
)

// Config defines library top level configuration.
type Config struct {
	ChainID        string                `mapstructure:"chain-id" json:"chain-id"`
	EncodingConfig params.EncodingConfig `mapstructure:"encoding-config" json:"encoding-config"`
	RootDir        string                `mapstructure:"root-dir" json:"root-dir"`
	RPCEndpoint    string                `mapstructure:"rpc-endpoint" json:"rpc-endpoint"`
}

// lib wide global singleton
var (
	libConfig  *Config
	sdkConfig  *sdk.Config
	initConfig sync.Once
)

// DefaultConfig returns library default configuration.
func DefaultConfig() *Config {
	return &Config{
		ChainID:        "planetmint-testnet-1",
		EncodingConfig: params.EncodingConfig{},
		RootDir:        "~/.planetmint-go/",
		RPCEndpoint:    "http://127.0.0.1:1317",
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		libConfig = DefaultConfig()
		sdkConfig = sdk.GetConfig()
	})
	return libConfig
}

// SetBech32PrefixForAccount sets the bech32 account prefix.
func (config *Config) SetBech32PrefixForAccount(bech32Prefix string) *Config {
	sdkConfig.SetBech32PrefixForAccount(bech32Prefix, "pub")
	return config
}

// SetEncodingConfig sets the encoding config and must not be nil.
func (config *Config) SetEncodingConfig(encodingConfig params.EncodingConfig) *Config {
	config.EncodingConfig = encodingConfig
	return config
}

// SetChainID sets the chain ID parameter.
func (config *Config) SetChainID(chainID string) *Config {
	config.ChainID = chainID
	return config
}

// SetRoot sets the root directory where to find the keyring.
func (config *Config) SetRoot(root string) *Config {
	config.RootDir = root
	return config
}

// SetRPCEndpoint sets the RPC endpoint to send requests to.
func (config *Config) SetRPCEndpoint(rpcEndpoint string) *Config {
	config.RPCEndpoint = rpcEndpoint
	return config
}
