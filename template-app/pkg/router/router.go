package router

import (
	"net/http"

	"github.com/gorilla/mux"
	promlib "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"{{.ServiceName}}/pkg/handler"
	"{{.ServiceName}}/pkg/middleware"
	"{{.ServiceName}}/pkg/prometheus"
)

const (
	metricsPath = "/metrics"
	healthPath  = "/health"
)

func New() http.Handler {
	promRegistry := promlib.NewRegistry()
	prom := prometheus.NewPrometheus(promRegistry)

	metrics := promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{})

	middlewares := []middleware.Middleware{
		middleware.HTTPMetrics(prom, "{{.ServiceName}}"),
		middleware.Recovery(prom, "{{.ServiceName}}"),
	}

	router := mux.NewRouter()

	router.Handle(healthPath, middleware.Wrap(handler.Health(), middlewares...)).Methods(http.MethodGet)
	router.Handle(metricsPath, middleware.Wrap(metrics, middlewares...)).Methods(http.MethodGet)

	return router
}
