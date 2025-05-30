package usecaseimpls

import (
	"github.com/eduardo-gualberto/go.git/internal/application/repositories/userrepo"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	"github.com/eduardo-gualberto/go.git/internal/core/usecases"
)

// ListUsersImpl implements the ListUsers use case.
type ListUsersImpl struct {
	repo userrepo.UserRepo
}

// NewListUsersImpl constructs a new ListUsersImpl with the given UserRepo.
func NewListUsersImpl(r userrepo.UserRepo) usecases.ListUsers {
	return &ListUsersImpl{repo: r}
}

// Execute retrieves all users.
func (u *ListUsersImpl) Execute() ([]*entities.UserEntity, error) {
	return u.repo.List()
}
