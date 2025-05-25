package usecases

import (
	"github.com/eduardo-gualberto/go.git/internal/core/interfaces"
)

type RespondToUser interface {
	Execute(interfaces.MessageReader) error
}
