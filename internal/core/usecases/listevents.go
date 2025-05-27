package usecases

import "github.com/eduardo-gualberto/go.git/internal/core/entities"

// ListEvents defines a use case for retrieving all events.
type ListEvents interface {
	// Execute retrieves all events.
	Execute() ([]*entities.EventEntity, error)
}
