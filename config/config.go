package config

import (
	"encoding/json"
	"sync"
)

const DefaultConfigTemplate = `
###############################################################################
###                         Planetmint                                      ###
###############################################################################

[planetmint]
validator-address = "{{ .PlmntConfig.ValidatorAddress }}"
mqtt-domain = "{{ .PlmntConfig.MqttDomain }}"
mqtt-port = {{ .PlmntConfig.MqttPort }}
mqtt-user = "{{ .PlmntConfig.MqttUser }}"
mqtt-password = "{{ .PlmntConfig.MqttPassword }}"
claim-host = "{{ .PlmntConfig.ClaimHost }}"
mqtt-tls = {{ .PlmntConfig.MqttTLS }}
issuer-host = "{{ .PlmntConfig.IssuerHost }}"
certs-path = "{{ .PlmntConfig.CertsPath }}"
`

// Config defines Planetmint's top level configuration
type Config struct {
	ValidatorAddress string `json:"validator-address" mapstructure:"validator-address"`
	MqttDomain       string `json:"mqtt-domain"       mapstructure:"mqtt-domain"`
	MqttPort         int    `json:"mqtt-port"         mapstructure:"mqtt-port"`
	MqttUser         string `json:"mqtt-user"         mapstructure:"mqtt-user"`
	MqttPassword     string `json:"mqtt-password"     mapstructure:"mqtt-password"`
	ClaimHost        string `json:"claim-host"        mapstructure:"claim-host"`
	MqttTLS          bool   `json:"mqtt-tls"          mapstructure:"mqtt-tls"`
	IssuerHost       string `json:"issuer-host"       mapstructure:"issuer-host"`
	CertsPath        string `json:"certs-path"        mapstructure:"certs-path"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		ValidatorAddress: "plmnt1w5dww335zhh98pzv783hqre355ck3u4w4hjxcx",
		MqttDomain:       "testnet-mqtt.rddl.io",
		MqttPort:         1886,
		MqttUser:         "user",
		MqttPassword:     "password",
		ClaimHost:        "https://testnet-p2r.rddl.io",
		MqttTLS:          true,
		IssuerHost:       "https://testnet-issuer.rddl.io",
		CertsPath:        "./certs/",
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
