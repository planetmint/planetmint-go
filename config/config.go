package app

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
watchmen-endpoint = "{{ .PlmntConfig.WatchmenConfig.Endpoint }}"
watchmen-port = {{ .PlmntConfig.WatchmenConfig.Port }}
`

// Config defines Planetmint's top level configuration
type Config struct {
	WatchmenConfig WatchmenConfig `mapstructure:"watchmen-config" json:"watchmen-config"`
}

// WatchmenConfig defines Planetmint's watchmen configuration
type WatchmenConfig struct {
	Endpoint string `mapstructure:"watchmen-endpoint" json:"watchmen-endpoint"`
	Port     int    `mapstructure:"watchmen-port" json:"watchmen-port"`
}

// cosmos-sdk wide global singleton
var (
	plmntConfig *Config
	initConfig  sync.Once
)

// DefaultConfig returns planetmint's default configuration.
func DefaultConfig() *Config {
	return &Config{
		WatchmenConfig: WatchmenConfig{
			Endpoint: "lab.r3c.network",
			Port:     7401,
		},
	}
}

// GetConfig returns the config instance for the SDK.
func GetConfig() *Config {
	initConfig.Do(func() {
		plmntConfig = DefaultConfig()
	})
	return plmntConfig
}

// SetWatchmenConfig sets Planetmint's watchmen configuration
func (config *Config) SetWatchmenConfig(watchmenConfig interface{}) {
	jsonWatchmenConfig, err := json.Marshal(watchmenConfig)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonWatchmenConfig, &config.WatchmenConfig)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", config.WatchmenConfig.Port)
}
