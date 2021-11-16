package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"{{.ServiceName}}/pkg/handler"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err, "Unexpected create request error")

	rr := httptest.NewRecorder()

	handler.Health().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Incorrect http status code")
}
