package utils

import (
	"errors"

	user "github.com/abyanmajid/codemore.io/services/user/proto"

	"github.com/abyanmajid/codemore.io/services/user/internal/database"
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

func UpdateUserFields(existingUser database.User, req *user.UpdateUserByIdRequest) (database.UpdateUserByIdParams, error) {
	username := existingUser.Username
	if req.GetUsername() != "" {
		username = req.GetUsername()
	}

	email := existingUser.Email
	if req.GetEmail() != "" {
		email = req.GetEmail()
	}

	password := existingUser.Password
	if req.GetPassword() != "" {
		hashedPassword, err := HashPassword(req.GetPassword())
		if err != nil {
			return database.UpdateUserByIdParams{}, errors.New("failed to hash new password: " + err.Error())
		}
		password = hashedPassword
	}

	roles := existingUser.Roles
	if len(req.GetRoles()) > 0 {
		roles = req.GetRoles()
	}

	xp := existingUser.Xp
	if req.GetXp() != 0 {
		xp = req.GetXp()
	}

	isBanned := existingUser.IsBanned
	if req.GetIsBanned() != existingUser.IsBanned {
		isBanned = req.GetIsBanned()
	}

	return database.UpdateUserByIdParams{
		ID:       existingUser.ID,
		Username: username,
		Email:    email,
		Password: password,
		Roles:    roles,
		Xp:       xp,
		IsBanned: isBanned,
	}, nil
}
