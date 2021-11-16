package config

import (
	"fmt"

	"{{.ServiceName}}/pkg/env"
)

const (
	appPortConfKey         = "APP_PORT"
	logLevelConfKey        = "LOG_LEVEL"
	serverTimeoutMSConfKey = "SERVER_TIMEOUT_MS"

	defaultLogLevel        = "INFO"
	defaultServerTimeoutMS = 10000
)

type Config struct {
	AppPort       string
	LogLevel      string
	ServerTimeOut int
}

func New() (*Config, error) {
	vars := &env.Vars{}

	appPort := vars.Mandatory(appPortConfKey)
	logLevel := vars.Optional(logLevelConfKey, defaultLogLevel)
	serverTimeoutMS := vars.OptionalInt(serverTimeoutMSConfKey, defaultServerTimeoutMS)

	if err := vars.Error(); err != nil {
		return nil, fmt.Errorf("config: environment variables: %w", err)
	}

	return &Config{
		AppPort:       appPort,
		LogLevel:      logLevel,
		ServerTimeOut: serverTimeoutMS,
	}, nil
}
