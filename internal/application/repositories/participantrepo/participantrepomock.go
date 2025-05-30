package participantrepo

import (
	"errors"

	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

type MockedParticipantRepo struct {
	participants []*entities.ParticipantEntity
}

func (m *MockedParticipantRepo) List() ([]*entities.ParticipantEntity, error) {
	return m.participants, nil
}

func (m *MockedParticipantRepo) GetByID(id int) (*entities.ParticipantEntity, error) {
	if id < len(m.participants) {
		return m.participants[id], nil
	}
	return nil, errors.New("participant not found")
}

func (m *MockedParticipantRepo) Create(e entities.ParticipantEntity) (*entities.ParticipantEntity, error) {
	m.participants = append(m.participants, &e)
	return &e, nil
}

func NewMockedParticipantRepo() *MockedParticipantRepo {
	return &MockedParticipantRepo{
		participants: make([]*entities.ParticipantEntity, 0),
	}
}