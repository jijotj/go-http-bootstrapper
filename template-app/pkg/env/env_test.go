package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"{{.ServiceName}}/pkg/env"
)

func TestMandatory(t *testing.T) {
	err := os.Setenv("TEST_KEY", "some-value")
	require.NoError(t, err, "Unexpected set env error")
	vars := env.Vars{}

	val := vars.Mandatory("TEST_KEY")
	err = vars.Error()

	assert.Equal(t, "some-value", val, "Incorrect env value")
	assert.NoError(t, err, "Unexpected get env error")

	err = os.Unsetenv("TEST_KEY")
	require.NoError(t, err, "Unexpected unset env error")
}

func TestMandatoryWhenEnvVarNotSet(t *testing.T) {
	vars := env.Vars{}

	val := vars.Mandatory("TEST_KEY")
	err := vars.Error()

	assert.Equal(t, "", val, "Incorrect env value")
	assert.EqualError(t, err, "missing mandatory config(s): TEST_KEY", "Incorrect error for missing mandatory config")
}

func TestOptional(t *testing.T) {
	err := os.Setenv("TEST_KEY", "some-value")
	require.NoError(t, err, "Unexpected set env error")
	vars := env.Vars{}

	val := vars.Optional("TEST_KEY", "")
	err = vars.Error()

	assert.Equal(t, "some-value", val, "Incorrect env value")
	assert.NoError(t, err, "Unexpected get env error")

	err = os.Unsetenv("TEST_KEY")
	require.NoError(t, err, "Unexpected unset env error")
}

func TestOptionalWhenEnvVarNotSet(t *testing.T) {
	vars := env.Vars{}

	val := vars.Optional("TEST_KEY", "some-value")
	err := vars.Error()

	assert.Equal(t, "some-value", val, "Incorrect env value")
	assert.NoError(t, err, "Unexpected get env error")
}

func TestOptionalInt(t *testing.T) {
	err := os.Setenv("TEST_KEY", "10")
	require.NoError(t, err, "Unexpected set env error")
	vars := env.Vars{}

	val := vars.OptionalInt("TEST_KEY", 0)
	err = vars.Error()

	assert.Equal(t, 10, val, "Incorrect env value")
	assert.NoError(t, err, "Unexpected get env error")

	err = os.Unsetenv("TEST_KEY")
	require.NoError(t, err, "Unexpected unset env error")
}

func TestOptionalIntWhenEnvVarNotSet(t *testing.T) {
	vars := env.Vars{}

	val := vars.OptionalInt("TEST_KEY", 10)
	err := vars.Error()

	assert.Equal(t, 10, val, "Incorrect env value")
	assert.NoError(t, err, "Unexpected get env error")
}

func TestOptionalIntWhenEnvVarIsMalformed(t *testing.T) {
	err := os.Setenv("TEST_KEY", "sample-value")
	require.NoError(t, err, "Unexpected set env error")
	vars := env.Vars{}

	val := vars.OptionalInt("TEST_KEY", 0)
	err = vars.Error()

	assert.Equal(t, 0, val, "Incorrect env value")
	assert.EqualError(t, err, "malformed config(s): TEST_KEY", "Incorrect error for malformed config")

	err = os.Unsetenv("TEST_KEY")
	require.NoError(t, err, "Unexpected unset env error")
}
