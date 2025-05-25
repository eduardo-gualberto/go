package messagereaders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/internal/application/models"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
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
		From:    r.m.GetContact().WaID,
		To:      r.m.GetMetadata().DisplayPhoneNumber,
		Kind:    entities.AudioContentKind,
		Content: "mensagem de documento affffs",
	}, nil
}

func NewDocumentMessageReader(wm *models.WebhookMessage) (interfaces.MessageReader, error) {
	return &DocumentMessageReader{m: wm}, nil
}
