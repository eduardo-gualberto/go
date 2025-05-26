package eventrepo

import (
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

// EventRepo defines the interface for event persistence operations.
type EventRepo interface {
	List() ([]*entities.EventEntity, error)
	Create(e entities.EventEntity) (*entities.EventEntity, error)
}