package metrics

import "time"

func (m *Metrics) ObserveKafkaPublish(
	startedAt time.Time,
	topic string,
	event string,
	err error,
) {
	duration := time.Since(startedAt)

	statusName := "success"
	if err != nil {
		statusName = "error"
	}

	m.kafkaPublishTotal.
		WithLabelValues(topic, event, statusName).
		Inc()

	m.kafkaPublishDuration.
		WithLabelValues(topic, event).
		Observe(duration.Seconds())
}
