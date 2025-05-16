package wabahandler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/eduardo-gualberto/go.git/core/interfaces"
	"github.com/eduardo-gualberto/go.git/core/usecases"
	"github.com/eduardo-gualberto/go.git/gateways/messagereaders"
	"github.com/eduardo-gualberto/go.git/gateways/models"
)

type WabaHandler struct {
	RespondToUser *usecases.RespondToUser
}

func (h *WabaHandler) HandleAuthenticate(w http.ResponseWriter, r *http.Request) {
	c := r.URL.Query().Get("hub.challenge")
	m := r.URL.Query().Get("hub.mode")
	v := r.URL.Query().Get("hub.verify_token")

	if m == "subscribe" && v == os.Getenv("META_VERIFY_TOKEN") {
		fmt.Fprint(w, c)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func (h *WabaHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Printf("POST /webhook: %s\n\n", body)
	fmt.Println("------------------------------------------------")
	msg, err := models.NewWebhookMessage(body)
	if err != nil {
		fmt.Printf("Error unserializing message: %v\n\n", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	msgType, err := msg.GetType()
	if err != nil {
		fmt.Printf("Error getting message type: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	var msgReader interfaces.MessageReader

	switch *msgType {
	case models.TextMessage:
		if msgReader, err = messagereaders.NewTextMessageReader(msg); err != nil {
			fmt.Printf("Error creating text msg reader: %v\n", err)
			return
		}
		fmt.Printf("Text message with: %s\n", msg.GetText().Body)
	case models.AudioMessage:
		if msgReader, err = messagereaders.NewAudioMessageReader(msg); err != nil {
			fmt.Printf("Error creating audio msg reader: %v\n", err)
			return
		}
		fmt.Printf("Audio message with: %s\n", msg.GetAudio().MimeType)
	case models.ImageMessage:
		if msgReader, err = messagereaders.NewImageMessageReader(msg); err != nil {
			fmt.Printf("Error creating image msg reader: %v\n", err)
			return
		}
		fmt.Printf("Image message with: %s\n", msg.GetImage().MimeType)
	case models.DocumentMessage:
		if msgReader, err = messagereaders.NewDocumentMessageReader(msg); err != nil {
			fmt.Printf("Error creating document msg reader: %v\n", err)
			return
		}
		fmt.Printf("Document message with: %s\n", msg.GetDocument().Filename)
	case models.StatusMessage:
		w.WriteHeader(http.StatusOK)
		fmt.Printf("Status message with: %s\n\n", msg.GetStatus())
		return
	}
	fmt.Println()

	result, err := h.RespondToUser.Execute(&usecases.RespondToUserInput{
		Reader: msgReader,
	})
	if err != nil {
		fmt.Printf("Error executing respond to user use case: %v\n", err)
		return
	}
	if result == nil {
		fmt.Println("respond to user use case returned result nil")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewWabaHandler(usecase *usecases.RespondToUser) *WabaHandler {
	return &WabaHandler{
		RespondToUser: usecase,
	}
}
