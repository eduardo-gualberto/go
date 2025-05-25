package chatcompletionllms

import (
	"context"

	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/openai/openai-go"
)

type OpenAiLLM struct {
}

func (llm *OpenAiLLM) Generate(msgs []*entities.MessageEntity) (*string, error) {
	client := openai.NewClient()
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{openai.UserMessage(msgs[0].Content)},
		Model:    openai.ChatModelGPT3_5Turbo,
	}
	completion, err := client.Chat.Completions.New(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return &completion.Choices[0].Message.Content, nil
}

func NewOpenAiLLM() *OpenAiLLM {
	return &OpenAiLLM{}
}
