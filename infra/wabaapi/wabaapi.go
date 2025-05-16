package wabaapi

import (
	"fmt"

	"github.com/eduardo-gualberto/go.git/infra/httpclient"
)

type WabaApi struct {
	client *httpclient.HttpClient
}

func NewWabaApi(client *httpclient.HttpClient) (*WabaApi, error) {
	// accessToken := os.Getenv("META_ACCESS_TOKEN")
	// baseUrl := os.Getenv("META_BASE_URL")
	// apiVersion := os.Getenv("META_API_VERSION")

	// client, err := httpclient.NewHttpClient(
	// 	httpclient.WithBaseUrl(fmt.Sprintf("%s/%s", baseUrl, apiVersion)),
	// 	httpclient.WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
	// 	httpclient.WithHeader("Content-Type", "application/json"),
	// )
	// if err != nil {
	// 	return nil, err
	// }
	return &WabaApi{client: client}, nil
}

func (w *WabaApi) sendMessage(from string, to string, messageType string, messageBody map[string]any) (bool, error) {
	// Construct the request body
	body := map[string]any{
		"messaging_product": "whatsapp",
		"to":                to,
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
