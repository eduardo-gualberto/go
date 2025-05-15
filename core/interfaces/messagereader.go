package interfaces

import "github.com/eduardo-gualberto/go.git/core/entities"

type MessageReader interface {
	Read() (*entities.MessageEntity, error)
}
