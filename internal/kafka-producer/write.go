package KafkaProducer

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

// TODO заменить event когда дело дойдет до записи в кафку, а так же создать схему в кафка ui
func (k *Kafka) WriteEvent(event string) error {
	payload, err := k.serializer.Serialize(
		k.topic,
		event,
	)
	if err != nil {
		return err
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	}

	deliveryChan := make(chan kafka.Event, 1)

	err = k.producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan

	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		return msg.TopicPartition.Error
	}

	return nil
}
