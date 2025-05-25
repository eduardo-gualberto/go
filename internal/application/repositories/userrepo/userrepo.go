package userrepo

import (
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
)

// UserRepo defines the interface for user persistence operations
type UserRepo interface {
	List() ([]*entities.UserEntity, error)
	GetByID(id int) (*entities.UserEntity, error)
	Create(e entities.UserEntity) (*entities.UserEntity, error)
}
