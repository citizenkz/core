package entity

import (
	"time"

	"github.com/citizenkz/core/ent"
)

type (
	User struct {
		ID        int       `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		BirthDate time.Time `json:"birth_date"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func MakeStorageUserToEntity(user *ent.User) *User {
	return &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
	}
}

func MakeStorageUserSliceToEntity(users []*ent.User) []*User {
	result := make([]*User, len(users))
	for _, user := range users {
		userEntity := MakeStorageUserToEntity(user)
		result = append(result, userEntity)
	}

	return result
}
