package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/participantrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// ListParticipantsImpl implements the ListParticipants use case.
type ListParticipantsImpl struct {
	repo participantrepo.ParticipantRepo
}

// NewListParticipantsImpl constructs a new ListParticipantsImpl with the given ParticipantRepo.
func NewListParticipantsImpl(r participantrepo.ParticipantRepo) usecases.ListParticipants {
	return &ListParticipantsImpl{repo: r}
}

// Execute retrieves all participants.
func (u *ListParticipantsImpl) Execute() ([]*entities.ParticipantEntity, error) {
	return u.repo.List()
}
