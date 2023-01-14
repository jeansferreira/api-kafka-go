package services

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaService struct {
	config *KafkaServiceConfig
}

type KafkaServiceConfig struct {
	Brokers []string
}

func NewKafkaService(config *KafkaServiceConfig) KafkaService {
	return KafkaService{
		config: config,
	}
}

func (k *KafkaService) Publish(topic string, messages []kafka.Message) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: k.config.Brokers,
		Topic:   topic,
	})

	return writer.WriteMessages(context.Background(), messages...)
}
