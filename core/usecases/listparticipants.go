package usecases

import "github.com/eduardo-gualberto/go.git/core/entities"

// ListParticipants defines a use case for retrieving all participants.
type ListParticipants interface {
	Execute() ([]*entities.ParticipantEntity, error)
}
