package config

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/planetmint/planetmint-go/lib"
	"github.com/planetmint/planetmint-go/lib/trustwallet"
	"github.com/rddl-network/go-utils/logger"
)

const DefaultConfigTemplate = `
###############################################################################
###                         Planetmint                                      ###
###############################################################################

[planetmint]
certs-path = "{{ .PlmntConfig.CertsPath }}"
claim-host = "{{ .PlmntConfig.ClaimHost }}"
issuer-host = "{{ .PlmntConfig.IssuerHost }}"
mqtt-domain = "{{ .PlmntConfig.MqttDomain }}"
mqtt-password = "{{ .PlmntConfig.MqttPassword }}"
mqtt-port = {{ .PlmntConfig.MqttPort }}
mqtt-tls = {{ .PlmntConfig.MqttTLS }}
mqtt-user = "{{ .PlmntConfig.MqttUser }}"
`

// ValAddr to be reomved see: https://github.com/planetmint/planetmint-go/issues/454
const ValAddr = "VALIDATOR_ADDRESS"

const (
	defaultCertsPath    = "./certs/"
	defaultClaimHost    = "https://testnet-p2r.rddl.io"
	defaultIssuerHost   = "https://testnet-issuer.rddl.io"
	defaultMqttDomain   = "testnet-mqtt.rddl.io"
	defaultMqttPassword = "password"
	defaultMqttPort     = 1886
	defaultMqttTLS      = true
	defaultMqttUser     = "user"
)

// Config defines Planetmint's top level configuration
type Config struct {
	CertsPath    string `json:"certs-path"    mapstructure:"certs-path"`
	ClaimHost    string `json:"claim-host"    mapstructure:"claim-host"`
	IssuerHost   string `json:"issuer-host"   mapstructure:"issuer-host"`
	MqttDomain   string `json:"mqtt-domain"   mapstructure:"mqtt-domain"`
	MqttPassword string `json:"mqtt-password" mapstructure:"mqtt-password"`
	MqttPort     int    `json:"mqtt-port"     mapstructure:"mqtt-port"`
	MqttTLS      bool   `json:"mqtt-tls"      mapstructure:"mqtt-tls"`
	MqttUser     string `json:"mqtt-user"     mapstructure:"mqtt-user"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		CertsPath:    defaultCertsPath,
		ClaimHost:    defaultClaimHost,
		IssuerHost:   defaultIssuerHost,
		MqttDomain:   defaultMqttDomain,
		MqttPassword: defaultMqttPassword,
		MqttPort:     defaultMqttPort,
		MqttTLS:      defaultMqttTLS,
		MqttUser:     defaultMqttUser,
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		plmntConfig = DefaultConfig()
	})
	return plmntConfig
}

// SetPlanetmintConfig sets Planetmint's configuration
func (config *Config) SetPlanetmintConfig(planetmintconfig interface{}) {
	jsonConfig, err := json.Marshal(planetmintconfig)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(jsonConfig, config); err != nil {
		panic(err)
	}
}

// GetNodeAddress retrieves the validator address through multiple methods
func (config *Config) GetNodeAddress() (address string) {
	var err error = nil
	// Check environment variable first <- this is used for test cases only
	if address = os.Getenv(ValAddr); address != "" {
		return
	}

	libConfig := lib.GetConfig()

	// Handle no Trust Wallet connected scenario
	if libConfig.GetSerialPort() == "" {
		address, err = getKeyringAddress(libConfig)
	} else { // Handle Trust Wallet connected scenario
		address, err = getTrustWalletAddress(libConfig)
	}

	if err != nil {
		msg := "Cannot get node address. Please configure a Trust Wallet or define at least one key pair in the utilized keyring."
		new_error := errors.New(msg + ": " + err.Error())
		panic(new_error)
	}
	return address
}

// getKeyringAddress retrieves the default validator address
func getKeyringAddress(libConfig *lib.Config) (address string, err error) {
	defaultRecord, err := libConfig.GetDefaultValidatorRecord()
	if err != nil {
		logger.GetLogger(logger.ERROR).Error("msg", err.Error())
		return
	}

	addr, err := defaultRecord.GetAddress()
	if err != nil {
		logger.GetLogger(logger.ERROR).Error("msg", err.Error())
		return
	}
	address = addr.String()
	return
}

// getTrustWalletAddress retrieves validator address from Trust Wallet
func getTrustWalletAddress(libConfig *lib.Config) (address string, err error) {
	connector, err := trustwallet.NewTrustWalletConnector(libConfig.GetSerialPort())
	if err != nil {
		logger.GetLogger(logger.ERROR).Error("msg", err.Error())
		return
	}

	keys, err := connector.GetPlanetmintKeys()
	if err != nil {
		logger.GetLogger(logger.ERROR).Error("msg", err.Error())
		return
	}
	address = keys.PlanetmintAddress
	return
}
