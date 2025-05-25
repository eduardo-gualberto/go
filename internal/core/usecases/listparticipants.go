package usecases

import "github.com/eduardo-gualberto/go.git/internal/core/entities"

// ListParticipants defines a use case for retrieving all participants.
type ListParticipants interface {
	Execute() ([]*entities.ParticipantEntity, error)
}
