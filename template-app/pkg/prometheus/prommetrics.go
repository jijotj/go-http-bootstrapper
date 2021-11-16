package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	Buckets = []float64{
		0.001, 0.002, 0.005, 0.010, 0.015, 0.020, 0.025, 0.030,
		0.035, 0.040, 0.05, 0.075, 0.1, 0.125, 0.150, 0.175,
		0.2, 0.25, 0.5, 1, 2.5, 5, 10,
	}

	HTTPHandlerLatencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "hs",
		Name:      "http_handler_latency",
		Help:      "HTTP handler latency histogram partitioned by different services and HTTP status code",
		Buckets:   Buckets,
	}, []string{"service", "path"})

	HTTPStatusCodeCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "hs",
		Name:      "http_status_code",
		Help:      "HTTP status codes partitioned by different services, API path, HTTP status code and error",
	}, []string{"service", "path", "code", "error"})

	PanicCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "hs",
		Name:      "panic_count",
		Help:      "count of server panics by service name",
	}, []string{"service"})
)
