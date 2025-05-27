package usecases

// CreateEvent defines a use case for creating an event.
type CreateEvent interface {
	// Execute creates a new event with the given participant ID, user ID, and recurrence rule.
	Execute(participantID int, userID int, rrule string) error
}
