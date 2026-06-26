package KafkaProducer

import "errors"

func (k *Kafka) IsInitialized() error {
	if k.serializer == nil {
		return errors.New("serializer is already closed")
	}

	if k.producer == nil {
		return errors.New("producer is already closed")
	}

	return nil
}

func (k *Kafka) HealthCheck() error {
	if k.producer == nil {
		return errors.New("producer is nil")
	}

	_, err := k.producer.GetMetadata(nil, false, 3000)
	return err
}
