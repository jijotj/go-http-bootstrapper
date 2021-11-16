package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"{{.ServiceName}}/pkg/lib"
	"{{.ServiceName}}/pkg/prometheus"
)

func Recovery(p *prometheus.Prometheus, serviceName string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					p.PanicCounter().WithLabelValues(serviceName).Inc()
					log.Error().Err(err.(error)).Msg("Panic recovery")
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					lib.WriteResponseJSON(w, map[string]string{"error": "There was an internal server error"})
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
