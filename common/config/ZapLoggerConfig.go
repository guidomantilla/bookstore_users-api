package config

import (
	"go.uber.org/zap"
)

var (
	zapLoggerConfigFlag *zapLoggerConfig
	ZapLogger           *zap.Logger
)

type zapLoggerConfig struct {
}

func ConfigZapLogger(environment string) {

	if zapLoggerConfigFlag == nil {

		var logConfig zap.Config
		if environment == "pro" {
			logConfig = zap.NewProductionConfig()
		} else {
			logConfig = zap.NewDevelopmentConfig()
		}

		var err error
		ZapLogger, err = logConfig.Build()
		if err != nil {
			panic(err)
		}

		zapLoggerConfigFlag = &zapLoggerConfig{}
	}
}
