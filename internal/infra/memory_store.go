package infra

import (
	"log"
	"sync"

	"github.com/krakatoa/learn-async-go/internal/domain"
)

type MemoryRepository struct {
	mu       sync.RWMutex
	messages []domain.Message
}

func NewMemoryRepository() domain.MessageRepository {
	return &MemoryRepository{
		messages: make([]domain.Message, 0),
	}
}

func (s *MemoryRepository) Save(message domain.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messages = append(s.messages, message)
	log.Printf("[Repo] Message saved in memory %+v", message.Text)
	return nil
}

func (s *MemoryRepository) GetAll() []domain.Message {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return append([]domain.Message(nil), s.messages...)
}
