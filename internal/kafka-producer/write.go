package KafkaProducer

import (
	"ab/internal/dto/kafka-messege-dto"
	"ab/pkg/shortcut"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func (k *Kafka) WriteEvent(event *kafkaMessageDto.UserInExperimentMessage) error {
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
		Key:   []byte(fmt.Sprintf("%d", event.UserID)),
		Value: payload,
	}

	deliveryChan := make(chan kafka.Event, 1)

	err = k.producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan

	msg, ok := e.(*kafka.Message)
	if !ok {
		return shortcut.ErrTypeCast
	}

	if msg.TopicPartition.Error != nil {
		return msg.TopicPartition.Error
	}

	return nil
}
