package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	Registry *prometheus.Registry

	requestsTotal        *prometheus.CounterVec
	requestDuration      *prometheus.HistogramVec
	activeExperiments    *prometheus.GaugeVec
	cacheRefreshTime     prometheus.Histogram
	kafkaPublishTotal    *prometheus.CounterVec
	kafkaPublishDuration *prometheus.HistogramVec
}

func New() *Metrics {
	reg := prometheus.NewRegistry()

	m := &Metrics{
		Registry: reg,

		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ab_grpc_requests_total",
				Help: "Total gRPC requests",
			},
			[]string{"method", "status"},
		),

		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "ab_grpc_request_duration_seconds",
				Help:    "gRPC request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method"},
		),

		activeExperiments: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "ab_active_experiments",
				Help: "Current number of active experiments",
			},
			[]string{"type"},
		),

		cacheRefreshTime: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Name:    "ab_cache_refresh_duration_seconds",
				Help:    "Cache refresh duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
		),

		kafkaPublishTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ab_kafka_publish_total",
				Help: "Total Kafka publish attempts",
			},
			[]string{"topic", "event", "status"},
		),

		kafkaPublishDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "ab_kafka_publish_duration_seconds",
				Help:    "Kafka publish duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"topic", "event"},
		),
	}

	reg.MustRegister(
		m.requestsTotal,
		m.requestDuration,
		m.activeExperiments,
		m.cacheRefreshTime,
		m.kafkaPublishTotal,
		m.kafkaPublishDuration,
	)

	return m
}
