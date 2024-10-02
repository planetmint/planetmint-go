package config

import (
	"encoding/json"
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
mqtt-domain = "{{ .PlmntConfig.MqttDomain }}"
mqtt-port = {{ .PlmntConfig.MqttPort }}
mqtt-user = "{{ .PlmntConfig.MqttUser }}"
mqtt-password = "{{ .PlmntConfig.MqttPassword }}"
claim-host = "{{ .PlmntConfig.ClaimHost }}"
mqtt-tls = {{ .PlmntConfig.MqttTLS }}
issuer-host = "{{ .PlmntConfig.IssuerHost }}"
certs-path = "{{ .PlmntConfig.CertsPath }}"
`

// ValAddr to be reomved see: https://github.com/planetmint/planetmint-go/issues/454
const ValAddr = "VALIDATOR_ADDRESS"

// Config defines Planetmint's top level configuration
type Config struct {
	MqttDomain   string `json:"mqtt-domain"   mapstructure:"mqtt-domain"`
	MqttPort     int    `json:"mqtt-port"     mapstructure:"mqtt-port"`
	MqttUser     string `json:"mqtt-user"     mapstructure:"mqtt-user"`
	MqttPassword string `json:"mqtt-password" mapstructure:"mqtt-password"`
	ClaimHost    string `json:"claim-host"    mapstructure:"claim-host"`
	MqttTLS      bool   `json:"mqtt-tls"      mapstructure:"mqtt-tls"`
	IssuerHost   string `json:"issuer-host"   mapstructure:"issuer-host"`
	CertsPath    string `json:"certs-path"    mapstructure:"certs-path"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		MqttDomain:   "testnet-mqtt.rddl.io",
		MqttPort:     1886,
		MqttUser:     "user",
		MqttPassword: "password",
		ClaimHost:    "https://testnet-p2r.rddl.io",
		MqttTLS:      true,
		IssuerHost:   "https://testnet-issuer.rddl.io",
		CertsPath:    "./certs/",
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
	err = json.Unmarshal(jsonConfig, &config)
	if err != nil {
		panic(err)
	}
}

func (config *Config) GetValidatorAddress() string {
	// Case: testing
	if os.Getenv(ValAddr) != "" {
		return os.Getenv(ValAddr)
	}

	libConfig := lib.GetConfig()

	// Case: No Trust Wallet connected
	if libConfig.GetSerialPort() == "" {
		defaultRecord, err := libConfig.GetDefaultValidatorRecord()
		if err != nil {
			logger.GetLogger(logger.ERROR).Error("msg", err.Error())
			return ""
		}
		addr, err := defaultRecord.GetAddress()
		if err != nil {
			logger.GetLogger(logger.ERROR).Error("msg", err.Error())
			return ""
		}

		return addr.String()
	}

	// Case: Trust Wallet connected
	connector, err := trustwallet.NewTrustWalletConnector(libConfig.GetSerialPort())
	if err != nil {
		logger.GetLogger(logger.ERROR).Error("msg", err.Error())
		return ""
	}

	keys, err := connector.GetPlanetmintKeys()
	if err != nil {
		logger.GetLogger(logger.ERROR).Error("msg", err.Error())
		return ""
	}

	return keys.PlanetmintAddress
}
