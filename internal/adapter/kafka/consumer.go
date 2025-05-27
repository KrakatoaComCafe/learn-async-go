package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/krakatoa/learn-async-go/internal/app"
)

type KafkaConsumer struct {
	group    sarama.ConsumerGroup
	topic    string
	appLogic app.MessageService
}

func NewKafkaConsumer(appLogic app.MessageService) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := []string{os.Getenv("KAFKA_BROKER")}
	group, err := sarama.NewConsumerGroup(brokers, "my-group", config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		group:    group,
		topic:    "messages",
		appLogic: appLogic,
	}, nil
}

func (c *KafkaConsumer) Start(ctx context.Context) {
	handler := consumerGroupHandler{
		svc: c.appLogic,
	}
	go func() {
		for {
			if ctx.Err() != nil {
				log.Printf("[Kafka] Context canceled, stopping consumer")
				return
			}

			if err := c.group.Consume(ctx, []string{c.topic}, handler); err != nil {
				log.Printf("[Kafka] Error consuming message: %+v", err)
			}
		}
	}()
}
func (c *KafkaConsumer) Close() error {
	log.Println("[Kafka] Closing consumer group...")
	return c.group.Close()
}

type consumerGroupHandler struct {
	svc app.MessageService
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		text := string(msg.Value)
		log.Printf("[Kafka Consumer] Message received %s", text)

		if err := h.svc.SaveMessage(text); err != nil {
			log.Println(fmt.Errorf("[Kafka Consumer] Error to save message [%s: %w]", text, err))
			continue
		}

		log.Printf("[Kafka Consumer] Message save")
		sess.MarkMessage(msg, "")
	}
	return nil
}
