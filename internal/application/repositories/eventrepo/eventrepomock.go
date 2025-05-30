package eventrepo

import (
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

type MockedEventRepo struct {
	events []*entities.EventEntity
}

func (m *MockedEventRepo) List() ([]*entities.EventEntity, error) {
	return m.events, nil
}

func (m *MockedEventRepo) Create(e entities.EventEntity) (*entities.EventEntity, error) {
	m.events = append(m.events, &e)
	return &e, nil
}

func NewMockedEventRepo() *MockedEventRepo {
	return &MockedEventRepo{
		events: make([]*entities.EventEntity, 0),
	}
}