package app

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	WATCHMEN_ENDPOINT string
	WATCHMEN_PORT     uint32
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	filename := fmt.Sprintf("./%s_config.json", env)
	gonfig.GetConf(filename, &configuration)

	return configuration
}
