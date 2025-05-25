package messagereaders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/internal/application/models"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
)

type ImageMessageReader struct {
	m *models.WebhookMessage
}

func (r ImageMessageReader) Read() (*entities.MessageEntity, error) {
	if r.m == nil {
		return nil, fmt.Errorf("no WebhookMessage provided to ImageMessageReader")
	}

	if !r.m.IsImageMessage() || r.m.GetImage() == nil {
		return nil, fmt.Errorf("provided WebhookMessage is not of image message type")
	}

	return &entities.MessageEntity{
		From:    r.m.GetContact().WaID,
		To:      r.m.GetMetadata().DisplayPhoneNumber,
		Kind:    entities.AudioContentKind,
		Content: "mensagem de imagem affffs",
	}, nil
}

func NewImageMessageReader(wm *models.WebhookMessage) (interfaces.MessageReader, error) {
	return &ImageMessageReader{m: wm}, nil
}
