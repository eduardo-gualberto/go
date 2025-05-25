package interfaces

import "github.com/eduardo-gualberto/go.git/internal/core/entities"

type MessageReader interface {
	Read() (*entities.MessageEntity, error)
}
