package messagereaders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/core/interfaces"
	"github.com/eduardo-gualberto/go.git/gateways/models"
)

type DocumentMessageReader struct {
	m *models.WebhookMessage
}

func (r DocumentMessageReader) Read() (*entities.MessageEntity, error) {
	if r.m == nil {
		return nil, fmt.Errorf("no WebhookMessage provided to DocumentMessageReader")
	}

	if !r.m.IsDocumentMessage() || r.m.GetDocument() == nil {
		return nil, fmt.Errorf("provided WebhookMessage is not of document message type")
	}

	return &entities.MessageEntity{
		From:    r.m.GetContact().Profile.Name,
		To:      r.m.GetContact().Profile.Name,
		Kind:    entities.AudioContentKind,
		Content: "mensagem de documento affffs",
	}, nil
}

func NewDocumentMessageReader(wm *models.WebhookMessage) (interfaces.MessageReader, error) {
	return &DocumentMessageReader{m: wm}, nil
}
