package app_test

import (
	"testing"

	"github.com/krakatoa/learn-async-go/internal/app"
	"github.com/krakatoa/learn-async-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(message domain.Message) error {
	args := m.Called(message)
	return args.Error(0)
}
func (m *MockRepository) GetAll() []domain.Message {
	return []domain.Message(nil)
}

type MockMessagePublisher struct {
	mock.Mock
}

func (m *MockMessagePublisher) Publish(text string) error {
	args := m.Called(text)
	return args.Error(0)
}

func TestMessageService_SaveMessage(t *testing.T) {
	repo := new(MockRepository)
	publisher := new(MockMessagePublisher)
	service := app.NewMessageService(repo, publisher)

	text := "hello world"
	message := domain.Message{
		ID:   "123",
		Text: text,
	}

	repo.On("Save", mock.MatchedBy(func(messageDomain domain.Message) bool {
		return messageDomain.Text == message.Text
	})).Return(nil)

	err := service.SaveMessage(text)
	assert.NoError(t, err)

	repo.AssertCalled(t, "Save", mock.MatchedBy(func(messageDomain domain.Message) bool {
		return messageDomain.Text == text
	}))
	publisher.AssertNotCalled(t, "Publish", text)
}

func TestMessageService_SendAndStoreMessage(t *testing.T) {
	repo := new(MockRepository)
	publisher := new(MockMessagePublisher)
	service := app.NewMessageService(repo, publisher)

	text := "hello, world!"

	publisher.On("Publish", text).Return(nil)

	err := service.SendAndStoreMessage(text)
	assert.NoError(t, err)

	publisher.AssertCalled(t, "Publish", text)
}
