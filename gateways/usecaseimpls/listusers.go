package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/core/entities"
	"github.com/eduardo-gualberto/go.git/gateways/repositories/userrepo"
)

// ListUsersImpl implements the ListUsers use case.
type ListUsersImpl struct {
	repo userrepo.UserRepo
}

// NewListUsersImpl constructs a new ListUsersImpl with the given UserRepo.
func NewListUsersImpl(r userrepo.UserRepo) *ListUsersImpl {
	return &ListUsersImpl{repo: r}
}

// Execute retrieves all users.
func (u *ListUsersImpl) Execute() ([]*entities.UserEntity, error) {
	return u.repo.List()
}
