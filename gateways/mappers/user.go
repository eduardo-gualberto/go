package mappers

import (
	"github.com/eduardo-gualberto/go.git/core/entities"

	db "github.com/eduardo-gualberto/go.git/infra/db/gen"
)

// UserDb2Entity maps a db.User to a domain UserEntity.
func UserDb2Entity(u *db.User) *entities.UserEntity {
	return &entities.UserEntity{
		ID:     int(u.ID),
		Number: u.UserNumber,
		Name:   u.UserName,
	}
}

// UserListDb2Entity maps a slice of db.User to a slice of domain UserEntity pointers.
func UserListDb2Entity(us []db.User) []*entities.UserEntity {
	out := make([]*entities.UserEntity, len(us))
	for _, u := range us {
		out = append(out, UserDb2Entity(&u))
	}
	return out
}

// UserEntity2Db maps a domain UserEntity to a db.User.
func UserEntity2Db(u *entities.UserEntity) *db.User {
	return &db.User{
		ID:         int64(u.ID),
		UserNumber: u.Number,
		UserName:   u.Name,
	}
}

// UserListEntity2Db maps a slice of domain UserEntity to a slice of db.User pointers.
func UserListEntity2Db(us []entities.UserEntity) []*db.User {
	out := make([]*db.User, len(us))
	for _, u := range us {
		out = append(out, UserEntity2Db(&u))
	}
	return out
}
