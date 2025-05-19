package domain

type MessageRepository interface {
	Save(msg Message) error
	GetAll() []Message
}
