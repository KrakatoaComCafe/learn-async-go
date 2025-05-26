package app

import (
	"github.com/google/uuid"
	"github.com/krakatoa/learn-async-go/internal/domain"
)

type MessageService interface {
	SaveMessage(text string) error
	SendAndStoreMessage(text string) error
	GetMessages() []domain.Message
}

type messageService struct {
	repo      domain.MessageRepository
	publisher MessagePublisher
}

type MessagePublisher interface {
	Publish(text string) error
}

func NewMessageService(repo domain.MessageRepository, publisher MessagePublisher) MessageService {
	return &messageService{
		repo:      repo,
		publisher: publisher,
	}
}

func (s *messageService) SendAndStoreMessage(text string) error {
	if err := s.publisher.Publish(text); err != nil {
		return err
	}
	return nil
}

func (s *messageService) SaveMessage(text string) error {
	msg := domain.Message{
		ID:   uuid.NewString(),
		Text: text,
	}
	return s.repo.Save(msg)
}

func (s *messageService) GetMessages() []domain.Message {
	return s.repo.GetAll()
}
