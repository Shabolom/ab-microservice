package KafkaProducer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/jsonschema"
)

type metrics interface {
	ObserveKafkaPublish(startedAt time.Time, topic string, event string, err error)
}
type Kafka struct {
	producer   *kafka.Producer
	serializer *jsonschema.Serializer
	metrics    metrics
	topic      string
}

func New(producer *kafka.Producer, serializer *jsonschema.Serializer, metrics metrics, topic string) *Kafka {
	return &Kafka{
		producer:   producer,
		serializer: serializer,
		metrics:    metrics,
		topic:      topic,
	}
}
