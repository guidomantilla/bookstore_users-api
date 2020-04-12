package config

import (
	"os"
	"strings"
)

var (
	propertiesConfigFlag *propertiesConfig
	Properties           map[string]string
)

type propertiesConfig struct {
}

func ConfigProperties() {

	if propertiesConfigFlag == nil {

		Properties = make(map[string]string, 0)
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			Properties[pair[0]] = pair[1]
		}
		propertiesConfigFlag = &propertiesConfig{}
	}
}
