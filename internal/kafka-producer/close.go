package KafkaProducer

func (k *Kafka) Close() error {
	if k.serializer != nil {
		if err := k.serializer.Close(); err != nil {
			return err
		}
	}

	if k.producer != nil && !k.producer.IsClosed() {
		k.producer.Close()
	}

	return nil
}
