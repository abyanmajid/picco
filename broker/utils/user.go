package utils

import (
	"time"

	"github.com/abyanmajid/codemore.io/broker/proto/user"
)

type User struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Roles     []string  `json:"roles"`
	Xp        int32     `json:"xp"`
	IsBanned  bool      `json:"is_banned"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DecodeProtoUser(u *user.User) User {
	return User{
		Id:        u.Id,
		Username:  u.Username,
		Email:     u.Email,
		Roles:     u.Roles,
		Xp:        u.Xp,
		IsBanned:  u.IsBanned,
		CreatedAt: u.CreatedAt.AsTime(),
		UpdatedAt: u.UpdatedAt.AsTime(),
	}
}

func DecodeMultipleProtoUsers(users []*user.User) []User {
	internalUsers := make([]User, len(users))
	for i, u := range users {
		internalUsers[i] = DecodeProtoUser(u)
	}
	return internalUsers
}
