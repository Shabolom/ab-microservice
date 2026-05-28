package KafkaProducer

func (k *Kafka) Close() {
	k.producer.Close()
}
