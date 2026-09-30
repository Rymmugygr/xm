package kafka

import (
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Kafka struct {
	brokers []string
	conn    *kgo.Client
}

func New(brokers []string) (*Kafka, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.AllowAutoTopicCreation(), // It may be reasonable to create all topics in advance
	)
	if err != nil {
		return nil, fmt.Errorf("kafka - connection error: %w", err)
	}

	k := &Kafka{
		brokers: brokers,
		conn:    client,
	}

	return k, nil
}

func (k *Kafka) Close() {
	if k.conn != nil {
		k.conn.Close()
	}
}
