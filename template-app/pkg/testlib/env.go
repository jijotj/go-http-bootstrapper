package testlib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func SetEnvVars(t *testing.T, configs map[string]string) func() {
	for key, value := range configs {
		err := os.Setenv(key, value)
		require.NoError(t, err, "Unexpected set env error")
	}

	return func() {
		for key := range configs {
			err := os.Unsetenv(key)
			require.NoError(t, err, "Unexpected set env error")
		}
	}
}
