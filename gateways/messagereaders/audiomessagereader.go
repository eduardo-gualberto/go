package messagereaders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/core/interfaces"
	"github.com/eduardo-gualberto/go.git/gateways/models"
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
		From:    r.m.GetContact().Profile.Name,
		To:      r.m.GetContact().Profile.Name,
		Kind:    entities.AudioContentKind,
		Content: "mensagem de áudio affffs",
	}, nil
}

func NewAudioMessageReader(wm *models.WebhookMessage) (interfaces.MessageReader, error) {
	return &AudioMessageReader{m: wm}, nil
}
