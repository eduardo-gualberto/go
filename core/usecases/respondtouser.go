package usecases

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/core/interfaces"
)

type RespondToUser struct {
	Llm    interfaces.ChatCompletionLLM
	Sender interfaces.MessageSender
}
type RespondToUserInput struct {
	Reader interfaces.MessageReader
}
type RespondToUserOutput struct {
	Success bool
}

func (u *RespondToUser) Execute(input *RespondToUserInput) (*RespondToUserOutput, error) {
	output := &RespondToUserOutput{}
	m, err := input.Reader.Read()

	if err != nil {
		output.Success = false
		return output, err
	}
	if m == nil {
		output.Success = false
		return output, fmt.Errorf("message reader returned nil message")
	}

	generated, err := u.Llm.Generate([]*entities.MessageEntity{m})
	if err != nil {
		output.Success = false
		return output, err
	}
	if generated == nil {
		output.Success = false
		return output, fmt.Errorf("failed in chat completion step")
	}
	response_message := &entities.MessageEntity{
		From:    m.To,
		To:      m.From,
		Kind:    entities.TextContentKind,
		Content: *generated,
	}

	if err = u.Sender.Send(response_message); err != nil {
		output.Success = false
		return output, err
	}
	output.Success = true
	return output, nil
}
