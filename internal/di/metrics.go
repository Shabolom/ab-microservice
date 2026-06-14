package di

import "ab/internal/metrics"

func (d *DI) GetMetrics() *metrics.Metrics {
	if d.metrics != nil {
		return d.metrics
	}

	d.metrics = metrics.New()

	return d.metrics
}
