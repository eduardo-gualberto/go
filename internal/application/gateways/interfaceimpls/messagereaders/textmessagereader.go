package messagereaders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/internal/application/models"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
)

type TextMessageReader struct {
	m *models.WebhookMessage
}

func (r TextMessageReader) Read() (*entities.MessageEntity, error) {
	if r.m == nil {
		return nil, fmt.Errorf("no WebhookMessage provided to TextMessageReader")
	}

	if !r.m.IsTextMessage() || r.m.GetText() == nil {
		return nil, fmt.Errorf("provided WebhookMessage is not of text message type")
	}

	return &entities.MessageEntity{
		From:    r.m.GetContact().WaID,
		To:      r.m.GetMetadata().DisplayPhoneNumber,
		Kind:    entities.AudioContentKind,
		Content: r.m.GetText().Body,
	}, nil
}

func NewTextMessageReader(wm *models.WebhookMessage) (interfaces.MessageReader, error) {
	return &TextMessageReader{m: wm}, nil
}
