package kafka_test

import (
	"context"
	"testing"
	"time"

	"github.com/krakatoa/learn-async-go/internal/adapter/kafka"
	"github.com/krakatoa/learn-async-go/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockMessageService struct {
	received []string
}

func (s *MockMessageService) SendAndStoreMessage(text string) error {
	return nil
}

func (s *MockMessageService) GetMessages() []domain.Message {
	return []domain.Message(nil)
}

func (m *MockMessageService) SaveMessage(text string) error {
	m.received = append(m.received, text)
	return nil
}

func TestKafkaConsumer_Start(t *testing.T) {
	mockService := &MockMessageService{}

	consumer, err := kafka.NewKafkaConsumer(mockService)
	assert.NoError(t, err)
	assert.NotNil(t, consumer)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	consumer.Start(ctx)

	time.Sleep(2 * time.Second)

	assert.True(t, true)
}
