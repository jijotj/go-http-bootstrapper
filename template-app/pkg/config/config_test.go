package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"{{.ServiceName}}/pkg/config"
	"{{.ServiceName}}/pkg/testlib"
)

func TestNewConfig(t *testing.T) {
	cleanup := testlib.SetEnvVars(t, map[string]string{
		"APP_PORT":          "8080",
		"LOG_LEVEL":         "DEBUG",
		"SERVER_TIMEOUT_MS": "2000",
	})
	defer cleanup()

	appConfig, err := config.New()

	assert.NoError(t, err, "Unexpected create config error")
	assert.Equal(t, "8080", appConfig.AppPort, "Incorrect app port value")
	assert.Equal(t, "DEBUG", appConfig.LogLevel, "Incorrect log level value")
	assert.Equal(t, 2000, appConfig.ServerTimeOut, "Incorrect server timeout ms value")
}

func TestNewConfigWithoutOptionalValues(t *testing.T) {
	cleanup := testlib.SetEnvVars(t, map[string]string{
		"APP_PORT": "8080",
	})
	defer cleanup()

	appConfig, err := config.New()

	assert.NoError(t, err, "Unexpected create config error")
	assert.Equal(t, "8080", appConfig.AppPort, "Incorrect app port value")
	assert.Equal(t, "INFO", appConfig.LogLevel, "Incorrect log level value")
	assert.Equal(t, 10000, appConfig.ServerTimeOut, "Incorrect server timeout ms value")
}

func TestNewConfigWhenMissingMandatoryValues(t *testing.T) {
	_, err := config.New()

	assert.EqualError(t, err, "config: environment variables: missing mandatory config(s): APP_PORT", "Incorrect missing mandatory config error")
}

func TestNewConfigWhenMalformedValues(t *testing.T) {
	cleanup := testlib.SetEnvVars(t, map[string]string{
		"APP_PORT":          "8080",
		"LOG_LEVEL":         "DEBUG",
		"SERVER_TIMEOUT_MS": "abc",
	})
	defer cleanup()

	_, err := config.New()

	assert.EqualError(t, err, `config: environment variables: malformed config(s): SERVER_TIMEOUT_MS`, "Incorrect server timeout ms error")
}
