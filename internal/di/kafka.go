package di

import KafkaProducer "ab/internal/kafka-producer"

func (d *DI) GetKafka() *KafkaProducer.Kafka {
	if d.kafka != nil {
		return d.kafka
	}

	return KafkaProducer.New(d.NewKafkaProducer(), d.NewKafkaJSONSerializer(), d.Config().Kafka.Topic)
}
