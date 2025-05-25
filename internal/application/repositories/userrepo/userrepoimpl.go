package userrepo

import (
	"context"
	"fmt"

	"github.com/eduardo-gualberto/go.git/internal/application/mappers"
	"github.com/eduardo-gualberto/go.git/internal/core/entities"
	db "github.com/eduardo-gualberto/go.git/internal/infra/db/gen"
	"github.com/jackc/pgx/v5"
)

// UserRepoImpl implements the UserRepo interface using a Postgres database.
type UserRepoImpl struct {
	q *db.Queries
}

// List returns all users from the database
func (r *UserRepoImpl) List() ([]*entities.UserEntity, error) {
	users, err := r.q.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return mappers.UserListDb2Entity(users), nil
}

// GetByID retrieves a user by its ID
func (r *UserRepoImpl) GetByID(id int) (*entities.UserEntity, error) {
	user, err := r.q.GetUserByID(context.Background(), int64(id))
	if err != nil {
		return nil, err
	}
	return mappers.UserDb2Entity(&user), nil
}

// Create inserts a new user into the database
func (r *UserRepoImpl) Create(e entities.UserEntity) (*entities.UserEntity, error) {
	user, err := r.q.CreateUser(context.Background(), db.CreateUserParams{
		UserNumber: e.Number,
		UserName:   e.Name,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("created user: %s, %s\n", user.UserName, user.UserNumber)
	return mappers.UserDb2Entity(&user), nil
}

// NewUserRepoImpl creates a new UserRepo backed by Postgres
func NewUserRepoImpl(conn *pgx.Conn) UserRepo {
	return &UserRepoImpl{q: db.New(conn)}
}
