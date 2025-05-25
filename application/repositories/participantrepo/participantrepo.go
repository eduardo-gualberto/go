package participantrepo

import (
	"github.com/eduardo-gualberto/go.git/core/entities"
)

type ParticipantRepo interface {
	List() ([]*entities.ParticipantEntity, error)
	GetByID(id int) (*entities.ParticipantEntity, error)
	Create(e entities.ParticipantEntity) (*entities.ParticipantEntity, error)
}
