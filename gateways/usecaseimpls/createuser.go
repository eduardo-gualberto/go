package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/gateways/repositories/userrepo"
)

// CreateUserImpl implements the CreateUser use case.
type CreateUserImpl struct {
	repo userrepo.UserRepo
}

// NewCreateUserImpl constructs a new CreateUserImpl with the given UserRepo.
func NewCreateUserImpl(r userrepo.UserRepo) *CreateUserImpl {
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
