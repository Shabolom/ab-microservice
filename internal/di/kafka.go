package di

import KafkaProducer "ab/internal/kafka-producer"

func (d *DI) GetKafka() *KafkaProducer.Kafka {
	if d.kafka != nil {
		return d.kafka
	}

	d.kafka = KafkaProducer.New(
		d.NewKafkaProducer(),
		d.NewKafkaJSONSerializer(),
		d.GetMetrics(),
		d.Config().Kafka.Topic,
	)

	return d.kafka
}
