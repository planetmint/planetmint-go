package lib

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/lib/params"
)

// Config defines library top level configuration.
type Config struct {
	ChainID        string                `json:"chain-id"        mapstructure:"chain-id"`
	ClientCtx      client.Context        `json:"client-ctx"      mapstructure:"client-ctx"`
	EncodingConfig params.EncodingConfig `json:"encoding-config" mapstructure:"encoding-config"`
	FeeDenom       string                `json:"fee-denom"       mapstructure:"fee-denom"`
	RootDir        string                `json:"root-dir"        mapstructure:"root-dir"`
	RPCEndpoint    string                `json:"rpc-endpoint"    mapstructure:"rpc-endpoint"`
	TxGas          uint64                `json:"tx-gas"          mapstructure:"tx-gas"`
}

// lib wide global singleton
var (
	libConfig  *Config
	sdkConfig  *sdk.Config
	initConfig sync.Once
	changeLock sync.Mutex
)

// DefaultConfig returns library default configuration.
func DefaultConfig() *Config {
	return &Config{
		ChainID:        "planetmint-testnet-1",
		ClientCtx:      client.Context{},
		EncodingConfig: params.EncodingConfig{},
		FeeDenom:       "plmnt",
		RootDir:        "~/.planetmint-go/",
		RPCEndpoint:    "http://127.0.0.1:26657",
		TxGas:          200000,
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		changeLock.Lock()
		defer changeLock.Unlock()
		libConfig = DefaultConfig()
		sdkConfig = sdk.GetConfig()
		libConfig.SetBech32PrefixForAccount("plmnt")

		encodingConfig := MakeEncodingConfig()
		libConfig.SetEncodingConfig(encodingConfig)
	})
	return libConfig
}

// SetBech32PrefixForAccount sets the bech32 account prefix.
func (config *Config) SetBech32PrefixForAccount(bech32Prefix string) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	sdkConfig.SetBech32PrefixForAccount(bech32Prefix, "pub")
	return config
}

// SetEncodingConfig sets the encoding config and must not be nil.
func (config *Config) SetEncodingConfig(encodingConfig params.EncodingConfig) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.EncodingConfig = encodingConfig
	return config
}

// SetChainID sets the chain ID parameter.
func (config *Config) SetChainID(chainID string) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.ChainID = chainID
	return config
}

// SetClientCtx sets the client context parameter.
func (config *Config) SetClientCtx(clientCtx client.Context) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.ClientCtx = clientCtx
	return config
}

// SetFeeDenom sets the fee denominator parameter.
func (config *Config) SetFeeDenom(feeDenom string) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.FeeDenom = feeDenom
	return config
}

// SetRoot sets the root directory where to find the keyring.
func (config *Config) SetRoot(root string) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.RootDir = root
	return config
}

// SetRPCEndpoint sets the RPC endpoint to send requests to.
func (config *Config) SetRPCEndpoint(rpcEndpoint string) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.RPCEndpoint = rpcEndpoint
	return config
}

// SetTxGas sets the amount of Gas for the TX that is send to the network
func (config *Config) SetTxGas(txGas uint64) *Config {
	changeLock.Lock()
	defer changeLock.Unlock()
	config.TxGas = txGas
	return config
}
