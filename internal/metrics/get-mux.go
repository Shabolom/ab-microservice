package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (m *Metrics) MetricsMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(
		"/metrics",
		promhttp.HandlerFor(m.Registry, promhttp.HandlerOpts{}),
	)

	return mux
}
