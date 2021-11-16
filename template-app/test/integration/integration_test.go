package integration_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	serverURL = "http://app:80"
)

func TestHealthCheck(t *testing.T) {
	res, err := http.Get(serverURL + "/health")
	require.NoError(t, err, "Unexpected create request error")

	assert.Equal(t, http.StatusOK, res.StatusCode, "Incorrect http status code")
}
