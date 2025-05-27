package kafka

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer() (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	brokers := []string{os.Getenv("KAFKA_BROKER")}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
		topic:    "messages",
	}, nil
}

func (p *KafkaProducer) Publish(text string) error {
	log.Printf("[Kafka Producer] Publicando mensagem: %s", text)

	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(text),
	}

	_, _, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("[Kafka Producer] Error to publish message %s: %v", p.topic, err)
		return fmt.Errorf("error to publish message into Kafka: %w", err)
	}

	log.Printf("[Kafka Producer] Published message %+v", msg.Value)
	return nil
}
