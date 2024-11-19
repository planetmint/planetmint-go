package lib

import (
	"errors"
	"os"
	"sync"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/lib/params"
)

var (
	// ErrInvalidConfig is returned when configuration validation fails
	ErrInvalidConfig = errors.New("invalid configuration")

	// Global singleton instances
	instance  *Config
	mu        sync.RWMutex
	once      sync.Once
	sdkConfig *sdk.Config
)

// Config defines the top-level configuration for the Planetmint library.
// All fields are exported to allow external access while maintaining
// thread-safe modifications through methods.
type Config struct {
	ChainID        string
	ClientCtx      client.Context
	EncodingConfig params.EncodingConfig
	FeeDenom       string
	RPCEndpoint    string
	RootDir        string
	SerialPort     string
	TxGas          uint64
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	return &Config{
		ChainID:        "planetmint-testnet-1",
		ClientCtx:      client.Context{},
		EncodingConfig: params.EncodingConfig{},
		FeeDenom:       "plmnt",
		RPCEndpoint:    "http://127.0.0.1:26657",
		RootDir:        "~/.planetmint-go/",
		TxGas:          200000,
	}
}

// GetConfig returns the singleton Config instance, initializing it if necessary.
func GetConfig() *Config {
	once.Do(func() {
		instance = NewConfig()
		sdkConfig = sdk.GetConfig()

		// Initialize default configuration
		instance.SetBech32PrefixForAccount("plmnt")
		encodingConfig := MakeEncodingConfig()
		instance.SetEncodingConfig(encodingConfig)
	})
	return instance
}

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	mu.RLock()
	defer mu.RUnlock()

	if c.ChainID == "" {
		return errors.New("chain ID cannot be empty")
	}
	if c.RPCEndpoint == "" {
		return errors.New("RPC endpoint cannot be empty")
	}
	if c.TxGas == 0 {
		return errors.New("transaction gas cannot be zero")
	}
	return nil
}

// Builder methods

func (c *Config) SetBech32PrefixForAccount(prefix string) *Config {
	mu.Lock()
	defer mu.Unlock()
	sdkConfig.SetBech32PrefixForAccount(prefix, "pub")
	return c
}

func (c *Config) SetEncodingConfig(config params.EncodingConfig) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.EncodingConfig = config
	return c
}

func (c *Config) SetChainID(chainID string) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.ChainID = chainID
	return c
}

func (c *Config) SetClientCtx(ctx client.Context) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.ClientCtx = ctx
	return c
}

func (c *Config) SetFeeDenom(denom string) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.FeeDenom = denom
	return c
}

func (c *Config) SetRoot(root string) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.RootDir = root
	return c
}

func (c *Config) SetRPCEndpoint(endpoint string) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.RPCEndpoint = endpoint
	return c
}

func (c *Config) SetTxGas(gas uint64) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.TxGas = gas
	return c
}

func (c *Config) SetSerialPort(port string) *Config {
	mu.Lock()
	defer mu.Unlock()
	c.SerialPort = port
	return c
}

// Getter methods

func (c *Config) GetSerialPort() string {
	mu.RLock()
	defer mu.RUnlock()
	return c.SerialPort
}

// Keyring operations

// GetLibKeyring returns a new keyring instance configured with the current settings.
func (c *Config) GetLibKeyring() (keyring.Keyring, error) {
	mu.RLock()
	defer mu.RUnlock()

	return keyring.New(
		"lib",
		keyring.BackendTest,
		c.RootDir,
		os.Stdin,
		c.EncodingConfig.Marshaler,
		[]keyring.Option{}...,
	)
}

// GetDefaultValidatorRecord returns the first validator record from the keyring.
func (c *Config) GetDefaultValidatorRecord() (*keyring.Record, error) {
	keyring, err := c.GetLibKeyring()
	if err != nil {
		return nil, err
	}

	records, err := keyring.List()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, errors.New("no keyring records found")
	}

	return records[0], nil
}
