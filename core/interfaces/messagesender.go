package interfaces

import "github.com/eduardo-gualberto/go.git/core/entities"

type MessageSender interface {
	Send(m *entities.MessageEntity) error
}
