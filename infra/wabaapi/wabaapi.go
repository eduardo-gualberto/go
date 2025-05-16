package wabaapi

import (
	"fmt"
	"regexp"

	"github.com/eduardo-gualberto/go.git/infra/httpclient"
)

type WabaApi struct {
	client *httpclient.HttpClient
}

func NewWabaApi(client *httpclient.HttpClient) (*WabaApi, error) {
	return &WabaApi{client: client}, nil
}

func (w *WabaApi) sendMessage(from string, to string, messageType string, messageBody map[string]any) (bool, error) {
	number := to
	r := regexp.MustCompile(`^55(\d{2})(\d{8})$`)
	to_components := r.FindStringSubmatch(to)
	if len(to_components) > 0 {
		number = fmt.Sprintf("55%s9%s", to_components[1], to_components[2])
	}

	body := map[string]any{
		"messaging_product": "whatsapp",
		"to":                number,
		"type":              messageType,
	}
	if messageType == "text" {
		body["text"] = messageBody
	} else {
		body["template"] = messageBody
	}

	_, err := w.client.Post(fmt.Sprintf("/%s/messages", from), httpclient.WithBody(body))
	if err != nil {
		return false, err
	}
	return true, nil
}
