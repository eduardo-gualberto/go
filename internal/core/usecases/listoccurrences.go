package usecases

import "github.com/eduardo-gualberto/go.git/internal/core/entities"

// ListOccurrences defines a use case for retrieving all occurrences.
type ListOccurrences interface {
	// Execute retrieves all occurrences.
	Execute() ([]*entities.OccurrenceEntity, error)
}
