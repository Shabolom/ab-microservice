package metrics

import "time"

func (m *Metrics) ObserveGRPCRequest(now time.Time, fullMethod string, err error) {
	duration := time.Since(now)

	statusName := "success"
	if err != nil {
		statusName = "error"
	}

	m.requestsTotal.
		WithLabelValues(fullMethod, statusName).
		Inc()

	m.requestDuration.
		WithLabelValues(fullMethod).
		Observe(duration.Seconds())
}
