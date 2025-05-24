package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/gateways/repositories/participantrepo"
)

type CreateParticipantImpl struct {
	repo participantrepo.ParticipantRepo
}

func NewCreateParticipantImpl(r participantrepo.ParticipantRepo) *CreateParticipantImpl {
	return &CreateParticipantImpl{
		repo: r,
	}
}

func (u *CreateParticipantImpl) Execute(number string, name string, user int) error {
	participant := entities.ParticipantEntity{
		Number: number,
		Name:   name,
		UserID: user,
	}
	_, err := u.repo.Create(participant)
	return err
}
