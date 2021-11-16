package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	promlib "github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"{{.ServiceName}}/pkg/middleware"
	"{{.ServiceName}}/pkg/prometheus"
)

func TestRecovery(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err, "Unexpected create request error")

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(fmt.Errorf("panic error"))
	})
	registry := promlib.NewRegistry()
	prom := prometheus.NewPrometheus(registry)
	rr := httptest.NewRecorder()

	recoverMiddleware := middleware.Recovery(prom, "{{.ServiceName}}")

	recoverMiddleware(nextHandler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Incorrect http status code")
	assert.JSONEq(t, `{"error":"There was an internal server error"}`, rr.Body.String(), "Incorrect http response body")
}
