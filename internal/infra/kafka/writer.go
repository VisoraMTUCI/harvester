package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	w *kafka.Writer
}

func NewKafkaWriter(
	brokers []string,
	topic string,
	writeTimeout time.Duration,
) (*KafkaWriter, error) {
	w := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		WriteTimeout: writeTimeout,
		Balancer:     &kafka.Hash{},
	}

	return &KafkaWriter{w: w}, nil
}

func (x *KafkaWriter) WriteMessage(ctx context.Context, msg Message) error {
	return x.w.WriteMessages(ctx, kafka.Message{
		Key:   msg.Key,
		Value: msg.Value,
	})
}

func (x *KafkaWriter) Close() error {
	return x.w.Close()
}
