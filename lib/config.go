package lib

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"strings"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/planetmint-go/lib/params"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// Config defines library top level configuration.
type Config struct {
	ChainID        string                `mapstructure:"chain-id" json:"chain-id"`
	EncodingConfig params.EncodingConfig `mapstructure:"encoding-config" json:"encoding-config"`
	GRPCEndpoint   string                `mapstructure:"grpc-endpoint" json:"grpc-endpoint"`
	GRPCTLSCert    string                `mapstructure:"grpc-tls-cert" json:"grpc-tls-cert"`
	RootDir        string                `mapstructure:"root-dir" json:"root-dir"`
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
		GRPCEndpoint:   "127.0.0.1:9090",
		GRPCTLSCert:    "",
		RootDir:        "~/.planetmint-go/",
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		libConfig = DefaultConfig()
		sdkConfig = sdk.GetConfig()
		libConfig.SetBech32PrefixForAccount("plmnt")
	})
	return libConfig
}

func (config *Config) GetGRPCConn() (grpcConn *grpc.ClientConn, err error) {
	creds := insecure.NewCredentials()
	// Configure TLS for remote connection.
	if !strings.Contains(config.GRPCEndpoint, "127.0.0.1") && !strings.Contains(config.GRPCEndpoint, "localhost") {
		cert, err := os.ReadFile(config.GRPCTLSCert)
		if err != nil {
			return nil, err
		}
		certPool := x509.NewCertPool()
		ok := certPool.AppendCertsFromPEM(cert)
		if !ok {
			return nil, err
		}
		tlsConfig := &tls.Config{
			RootCAs: certPool,
		}
		creds = credentials.NewTLS(tlsConfig)
	}
	grpcConn, err = grpc.Dial(
		config.GRPCEndpoint,
		grpc.WithTransportCredentials(creds),
	)
	return
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

// SetGRPCEndpoint sets the gRPC endpoint to send requests to.
func (config *Config) SetGRPCEndpoint(grpcEndpoint string) *Config {
	config.GRPCEndpoint = grpcEndpoint
	return config
}

// SetGRPCTLSCert sets the gRPC TLS certificate to use for communication.
func (config *Config) SetGRPCTLSCert(grpcTLSCert string) *Config {
	config.GRPCTLSCert = grpcTLSCert
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
