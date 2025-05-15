package interfaces

import "github.com/eduardo-gualberto/go.git/core/entities"

type ChatCompletionLLM interface {
	Generate([]*entities.MessageEntity) (*string, error)
}
