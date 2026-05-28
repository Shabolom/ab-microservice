package di

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func (d *DI) NewProducer() *kafka.Writer {
	if d.kafkaProducer != nil {
		return d.kafkaProducer
	}

	producer := &kafka.Writer{
		Addr:                   kafka.TCP(d.Config().Kafka.Brokers...),
		Topic:                  d.Config().Kafka.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
		WriteTimeout:           10 * time.Second,
		ReadTimeout:            10 * time.Second,
		RequiredAcks:           kafka.RequireOne,
	}

	d.kafkaProducer = producer
	return producer
}
