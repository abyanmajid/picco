package utils

import (
	user "github.com/abyanmajid/codemore.io/user/proto"

	"github.com/abyanmajid/codemore.io/user/internal/database"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func HashPassword(password string) (string, error) {
	cost := 14 // You can adjust the cost factor, 14 is considered very secure but will be slower
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func EncodeProtoUser(u database.User) *user.User {
	return &user.User{
		Id:        u.ID.String(),
		Username:  u.Username,
		Email:     u.Email,
		Roles:     u.Roles,
		Xp:        u.Xp,
		IsBanned:  u.IsBanned,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}
