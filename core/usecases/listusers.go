package usecases

import "github.com/eduardo-gualberto/go.git/core/entities"

// ListUsers defines a use case for retrieving all users.
type ListUsers interface {
	Execute() ([]*entities.UserEntity, error)
}
