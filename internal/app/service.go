package app

import (
	"github.com/google/uuid"
	"github.com/krakatoa/learn-async-go/internal/domain"
)

type MessageService struct {
	repo      domain.MessageRepository
	publisher MessagePublisher
}

type MessagePublisher interface {
	Publish(text string) error
}

func NewMessageService(repo domain.MessageRepository, publisher MessagePublisher) *MessageService {
	return &MessageService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *MessageService) SendAndStoreMessage(text string) error {
	if err := s.publisher.Publish(text); err != nil {
		return err
	}
	return nil
}

func (s *MessageService) SaveMessage(text string) error {
	msg := domain.Message{
		ID:   uuid.NewString(),
		Text: text,
	}
	return s.repo.Save(msg)
}

func (s *MessageService) GetMessages() []domain.Message {
	return s.repo.GetAll()
}
