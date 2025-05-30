package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/userrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// CreateUserImpl implements the CreateUser use case.
type CreateUserImpl struct {
	repo userrepo.UserRepo
}

// NewCreateUserImpl constructs a new CreateUserImpl with the given UserRepo.
func NewCreateUserImpl(r userrepo.UserRepo) usecases.CreateUser {
	return &CreateUserImpl{repo: r}
}

// Execute creates a new user with the provided name and number.
func (u *CreateUserImpl) Execute(name string, number string) error {
	user := entities.UserEntity{
		Number: number,
		Name:   name,
	}
	_, err := u.repo.Create(user)
	return err
}
