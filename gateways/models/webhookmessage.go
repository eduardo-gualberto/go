package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type WebhookMessage struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Field string `json:"field"`
			Value struct {
				MessagingProduct string                    `json:"messaging_product"`
				Metadata         WebhookMessageMetadata    `json:"metadata"`
				Contacts         *[]WebhookMessageContact  `json:"contacts,omitempty"`
				Messages         *[]WebhookMessageMessages `json:"messages,omitempty"`
				Statuses         *[]struct {
					ID           string `json:"id"`
					Status       string `json:"status"`
					Timestamp    string `json:"timestamp"`
					RecipientID  string `json:"recipient_id"`
					Conversation struct {
						ID                  string `json:"id"`
						ExpirationTimestamp string `json:"expiration_timestamp"`
						Origin              struct {
							Type string `json:"type"`
						} `json:"origin"`
					} `json:"conversation"`
					Pricing struct {
						Billable     bool   `json:"billable"`
						PricingModel string `json:"pricing_model"`
						Category     string `json:"category"`
					} `json:"pricing"`
				} `json:"statuses,omitempty"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}

type WebhookMessageMetadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type WebhookMessageMessages struct {
	From      string                  `json:"from"`
	ID        string                  `json:"id"`
	Timestamp string                  `json:"timestamp"`
	Type      string                  `json:"type"`
	Context   *WebhookMessageContext  `json:"context,omitempty"`
	Text      *WebhookMessageText     `json:"text,omitempty"`
	Audio     *WebhookMessageAudio    `json:"audio,omitempty"`
	Image     *WebhookMessageImage    `json:"image,omitempty"`
	Document  *WebhookMessageDocument `json:"document,omitempty"`
}

type WebhookMessageContact struct {
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	WaID string `json:"wa_id"`
}

type WebhookMessageText struct {
	Body string `json:"body"`
}

type WebhookMessageAudio struct {
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
	Voice    bool   `json:"voice"`
}

type WebhookMessageImage struct {
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type WebhookMessageDocument struct {
	Filename string `json:"filename"`
	MimeType string `json:"mime_type"`
	Sha256   string `json:"sha256"`
	ID       string `json:"id"`
}

type WebhookMessageContext struct {
	From string `json:"from"`
	ID   string `json:"id"`
}

type WebhookMessageType string

var (
	TextMessage     WebhookMessageType = "text"
	AudioMessage    WebhookMessageType = "audio"
	ImageMessage    WebhookMessageType = "image"
	DocumentMessage WebhookMessageType = "document"
	StatusMessage   WebhookMessageType = "statuses"
)

func NewWebhookMessage(b []byte) (*WebhookMessage, error) {
	var data WebhookMessage
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (w WebhookMessage) GetType() (*WebhookMessageType, error) {
	if w.IsTextMessage() {
		return &TextMessage, nil
	}
	if w.IsAudioMessage() {
		return &AudioMessage, nil
	}
	if w.IsImageMessage() {
		return &ImageMessage, nil
	}
	if w.IsDocumentMessage() {
		return &DocumentMessage, nil
	}
	if w.IsStatusMessage() {
		return &StatusMessage, nil
	}
	return nil, fmt.Errorf("unknown webhook message type")
}

func (w WebhookMessage) IsStatusMessage() bool {
	return w.GetStatus() != nil
}

func (w WebhookMessage) IsReply() bool {
	return w.GetContext() != nil
}

func (w WebhookMessage) IsTextMessage() bool {
	return w.GetText() != nil
}

func (w WebhookMessage) IsAudioMessage() bool {
	return w.GetAudio() != nil
}

func (w WebhookMessage) IsImageMessage() bool {
	return w.GetImage() != nil
}
func (w WebhookMessage) IsDocumentMessage() bool {
	return w.GetDocument() != nil
}

func (w WebhookMessage) GetContext() *WebhookMessageContext {
	defer recover()
	if message := w.getFirstMessage(); message != nil {
		return message.Context
	}
	return nil
}

func (w WebhookMessage) GetMetadata() *WebhookMessageMetadata {
	defer recover()
	return &w.Entry[0].Changes[0].Value.Metadata
}

func (w WebhookMessage) GetContact() *WebhookMessageContact {
	defer recover()
	if contacts := w.Entry[0].Changes[0].Value.Contacts; contacts != nil {
		return &(*contacts)[0]
	}
	return nil
}

func (w WebhookMessage) GetText() *WebhookMessageText {
	defer recover()
	if message := w.getFirstMessage(); message != nil {
		return message.Text
	}
	return nil
}

func (w WebhookMessage) GetAudio() *WebhookMessageAudio {
	defer recover()
	if message := w.getFirstMessage(); message != nil {
		return message.Audio
	}
	return nil
}

func (w WebhookMessage) GetImage() *WebhookMessageImage {
	defer recover()
	if message := w.getFirstMessage(); message != nil {
		return message.Image
	}
	return nil
}

func (w WebhookMessage) GetDocument() *WebhookMessageDocument {
	defer recover()
	if message := w.getFirstMessage(); message != nil {
		return message.Document
	}
	return nil
}

func (w WebhookMessage) GetStatus() any {
	defer recover()
	if statuses := w.Entry[0].Changes[0].Value.Statuses; statuses != nil {
		return (*statuses)[0]
	}
	return nil
}

func (w WebhookMessage) getFirstMessage() *WebhookMessageMessages {
	messages := w.Entry[0].Changes[0].Value.Messages
	if messages == nil {
		return nil
	}
	if len(*messages) == 0 {
		return nil
	}
	return &(*messages)[0]
}
