package interfaces

import "github.com/eduardo-gualberto/go.git/internal/core/entities"

type ChatCompletionLLM interface {
	Generate([]*entities.MessageEntity) (*string, error)
}
