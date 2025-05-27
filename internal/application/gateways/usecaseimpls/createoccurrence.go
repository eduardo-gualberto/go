package usecaseimpls

import (
	"time"

	"github.com/eduardo-gualberto/go.git/internal/application/repositories/occurrencerepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// CreateOccurrenceImpl implements the CreateOccurrence use case.
type CreateOccurrenceImpl struct {
	repo occurrencerepo.OccurrenceRepo
}

// NewCreateOccurrenceImpl constructs a new CreateOccurrenceImpl with the given OccurrenceRepo.
func NewCreateOccurrenceImpl(r occurrencerepo.OccurrenceRepo) usecases.CreateOccurrence {
	return &CreateOccurrenceImpl{repo: r}
}

// Execute creates a new occurrence for the provided event ID at the specified time.
func (u *CreateOccurrenceImpl) Execute(eventID int, occursAt time.Time) error {
	occ := entities.OccurrenceEntity{
		EventID:  eventID,
		OccursAt: occursAt,
	}
	_, err := u.repo.Create(occ)
	return err
}
