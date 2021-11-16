package middleware

import (
	"net/http"
	"strconv"

	promlib "github.com/prometheus/client_golang/prometheus"

	"{{.ServiceName}}/pkg/lib"
	"{{.ServiceName}}/pkg/prometheus"
)

func HTTPMetrics(p *prometheus.Prometheus, serviceName string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			timer := promlib.NewTimer(p.HTTPHandlerLatencyHistogram().WithLabelValues(serviceName, path))
			defer timer.ObserveDuration()

			ww := lib.NewRecordingWriter(w)
			next.ServeHTTP(ww, r)

			p.HTTPStatusCodeCounter().WithLabelValues(serviceName, path, strconv.Itoa(ww.Status), ww.Err.Code).Inc()
		})
	}
}
