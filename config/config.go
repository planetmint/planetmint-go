package app

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	WATCHMEN_ENDPOINT string
	WATCHMEN_PORT     int
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	err := gonfig.GetConf(filePath, &configuration)

	if err != nil {
		panic(err)
	}

	return configuration
}
