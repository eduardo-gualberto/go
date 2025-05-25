package messagesenders

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
	"github.com/eduardo-gualberto/go.git/internal/infra/wabaapi"
)

type WabaMessageSender struct {
	Api *wabaapi.WabaApi
}

func (s *WabaMessageSender) Send(m *entities.MessageEntity) error {
	ok, err := s.Api.SendText(m.To, s.Api.WithBody(m.Content))
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("could not send text message")
	}
	return nil
}

func NewWabaMessageSender(api *wabaapi.WabaApi) (interfaces.MessageSender, error) {
	return &WabaMessageSender{
		Api: api,
	}, nil
}
