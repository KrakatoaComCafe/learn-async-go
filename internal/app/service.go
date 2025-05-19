package app

import (
	"github.com/google/uuid"
	"github.com/krakatoa/learn-async-go/internal/domain"
)

type MessageService struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) *MessageService {
	return &MessageService{
		repo: repo,
	}
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
