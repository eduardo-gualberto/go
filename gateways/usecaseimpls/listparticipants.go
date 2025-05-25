package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/gateways/repositories/participantrepo"
)

// ListParticipantsImpl implements the ListParticipants use case.
type ListParticipantsImpl struct {
	repo participantrepo.ParticipantRepo
}

// NewListParticipantsImpl constructs a new ListParticipantsImpl with the given ParticipantRepo.
func NewListParticipantsImpl(r participantrepo.ParticipantRepo) *ListParticipantsImpl {
	return &ListParticipantsImpl{repo: r}
}

// Execute retrieves all participants.
func (u *ListParticipantsImpl) Execute() ([]*entities.ParticipantEntity, error) {
	return u.repo.List()
}
