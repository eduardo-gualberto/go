package messagereaders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/internal/application/models"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
)

type AudioMessageReader struct {
	m *models.WebhookMessage
}

func (r AudioMessageReader) Read() (*entities.MessageEntity, error) {
	if r.m == nil {
		return nil, fmt.Errorf("no WebhookMessage provided to AudioMessageReader")
	}

	if !r.m.IsAudioMessage() || r.m.GetAudio() == nil {
		return nil, fmt.Errorf("provided WebhookMessage is not of audio message type")
	}

	return &entities.MessageEntity{
		From:    r.m.GetContact().WaID,
		To:      r.m.GetMetadata().DisplayPhoneNumber,
		Kind:    entities.AudioContentKind,
		Content: "mensagem de áudio affffs",
	}, nil
}

func NewAudioMessageReader(wm *models.WebhookMessage) (interfaces.MessageReader, error) {
	return &AudioMessageReader{m: wm}, nil
}
