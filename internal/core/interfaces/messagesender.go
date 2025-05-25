package interfaces

import "github.com/eduardo-gualberto/go.git/internal/core/entities"

type MessageSender interface {
	Send(m *entities.MessageEntity) error
}
