package app

import (
	"encoding/json"
	"sync"
)

const DefaultConfigTemplate = `
###############################################################################
###                         Planetmint                                      ###
###############################################################################

[planetmint]
osc-service-port = {{ .PlmntConfig.OSCServicePort }}
watchmen-endpoint = "{{ .PlmntConfig.WatchmenEndpoint }}"
watchmen-port = {{ .PlmntConfig.WatchmenPort }}
`

// Config defines Planetmint's top level configuration
type Config struct {
	OSCServicePort   int    `mapstructure:"osc-service-port" json:"osc-service-port"`
	WatchmenEndpoint string `mapstructure:"watchmen-endpoint" json:"watchmen-endpoint"`
	WatchmenPort     int    `mapstructure:"watchmen-port" json:"watchmen-port"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		OSCServicePort:   8766,
		WatchmenEndpoint: "lab.r3c.network",
		WatchmenPort:     7401,
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		plmntConfig = DefaultConfig()
	})
	return plmntConfig
}

// SetWatchmenConfig sets Planetmint's configuration
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
