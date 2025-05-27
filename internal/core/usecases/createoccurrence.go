package usecases

import "time"

// CreateOccurrence defines a use case for creating an occurrence.
type CreateOccurrence interface {
	// Execute creates a new occurrence for the given event ID at the specified time.
	Execute(eventID int, occursAt time.Time) error
}
