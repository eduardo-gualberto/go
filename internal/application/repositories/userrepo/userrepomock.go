package userrepo

import (
	"errors"

	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

type MockedUserRepo struct {
	users []*entities.UserEntity
}

func (m *MockedUserRepo) List() ([]*entities.UserEntity, error) {
	return m.users, nil
}

func (m *MockedUserRepo) GetByID(id int) (*entities.UserEntity, error) {
	if id < len(m.users) {
		return m.users[id], nil
	}
	return nil, errors.New("user not found")
}

func (m *MockedUserRepo) Create(e entities.UserEntity) (*entities.UserEntity, error) {
	m.users = append(m.users, &e)
	return &e, nil
}

func NewMockedUserRepo() *MockedUserRepo {
	return &MockedUserRepo{
		users: make([]*entities.UserEntity, 0),
	}
}
