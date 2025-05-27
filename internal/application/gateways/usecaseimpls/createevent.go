package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/eventrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// CreateEventImpl implements the CreateEvent use case.
type CreateEventImpl struct {
	repo eventrepo.EventRepo
}

// NewCreateEventImpl constructs a new CreateEventImpl with the given EventRepo.
func NewCreateEventImpl(r eventrepo.EventRepo) usecases.CreateEvent {
	return &CreateEventImpl{repo: r}
}

// Execute creates a new event with the provided participant ID, user ID, and recurrence rule.
func (u *CreateEventImpl) Execute(participantID int, userID int, rrule string) error {
	event := entities.EventEntity{
		ParticipantID: participantID,
		UserID:        userID,
		Rrule:         rrule,
	}
	_, err := u.repo.Create(event)
	return err
}
