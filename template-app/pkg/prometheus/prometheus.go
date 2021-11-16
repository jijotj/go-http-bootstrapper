package prometheus

import "github.com/prometheus/client_golang/prometheus"

type Prometheus struct {
	registry *prometheus.Registry
}

func NewPrometheus(registry *prometheus.Registry) *Prometheus {
	return &Prometheus{registry: registry}
}

func (p *Prometheus) HTTPHandlerLatencyHistogram() *prometheus.HistogramVec {
	return registerHistogram(p.registry, HTTPHandlerLatencyHistogram)
}

func (p *Prometheus) HTTPStatusCodeCounter() *prometheus.CounterVec {
	return registerCounter(p.registry, HTTPStatusCodeCounter)
}

func (p *Prometheus) PanicCounter() *prometheus.CounterVec {
	return registerCounter(p.registry, PanicCounter)
}

func registerHistogram(registry *prometheus.Registry, histogram *prometheus.HistogramVec) *prometheus.HistogramVec {
	if err := registry.Register(histogram); err != nil {
		if _, ok := err.(prometheus.AlreadyRegisteredError); !ok {
			panic(err)
		}
	}

	return histogram
}

func registerCounter(registry *prometheus.Registry, counter *prometheus.CounterVec) *prometheus.CounterVec {
	if err := registry.Register(counter); err != nil {
		if _, ok := err.(prometheus.AlreadyRegisteredError); !ok {
			panic(err)
		}
	}

	return counter
}
