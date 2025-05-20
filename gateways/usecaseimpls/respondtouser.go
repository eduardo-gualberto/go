package usecaseimpls

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/core/interfaces"
)

type RespondToUserImpl struct {
	Llm    interfaces.ChatCompletionLLM
	Sender interfaces.MessageSender
}

func (u RespondToUserImpl) Execute(reader interfaces.MessageReader) error {
	m, err := reader.Read()

	if err != nil {
		return err
	}
	if m == nil {
		return fmt.Errorf("message reader returned nil message")
	}

	generated, err := u.Llm.Generate([]*entities.MessageEntity{m})
	if err != nil {
		return err
	}
	if generated == nil {
		return fmt.Errorf("failed in chat completion step")
	}
	response_message := &entities.MessageEntity{
		From:    m.To,
		To:      m.From,
		Kind:    entities.TextContentKind,
		Content: *generated,
	}

	if err = u.Sender.Send(response_message); err != nil {
		return err
	}
	return nil
}

func NewRespondToUserImpl(llm interfaces.ChatCompletionLLM, sender interfaces.MessageSender) *RespondToUserImpl {
	return &RespondToUserImpl{
		Llm:    llm,
		Sender: sender,
	}
}
