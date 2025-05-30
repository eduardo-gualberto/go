package occurrencerepo

import (
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

type MockedOccurrenceRepo struct {
	occurrences []*entities.OccurrenceEntity
}

func (m *MockedOccurrenceRepo) List() ([]*entities.OccurrenceEntity, error) {
	return m.occurrences, nil
}

func (m *MockedOccurrenceRepo) Create(e entities.OccurrenceEntity) (*entities.OccurrenceEntity, error) {
	m.occurrences = append(m.occurrences, &e)
	return &e, nil
}

func NewMockedOccurrenceRepo() *MockedOccurrenceRepo {
	return &MockedOccurrenceRepo{
		occurrences: make([]*entities.OccurrenceEntity, 0),
	}
}