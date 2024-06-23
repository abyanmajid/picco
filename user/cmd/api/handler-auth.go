package main

import (
	"context"
	"time"

	user "github.com/abyanmajid/codemore.io/user/proto"
	"github.com/abyanmajid/codemore.io/user/utils"
	"github.com/golang-jwt/jwt/v4"
)

func (s *UserServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	requestPayload := LoginPayload{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// Check if email exists
	u, err := s.DB.GetUserByEmail(ctx, requestPayload.Email)
	if err != nil {
		return nil, err
	}

	// Check if password matches
	err = utils.CheckPassword(requestPayload.Password, u.Password)
	if err != nil {
		return nil, err
	}

	// Create JWT token
	var token Token
	err = token.Create(u.Email, SECRET_KEY)
	if err != nil {
		return nil, err
	}

	res := &user.LoginResponse{
		Token: token.Value,
	}

	return res, nil
}

func (s *UserServer) Refresh(ctx context.Context, req *user.RefreshRequest) (*user.RefreshResponse, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		req.GetRefreshToken(),
		claims,
		func(token *jwt.Token) (any, error) {
			return SECRET_KEY, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	// Issue new token if the old token is within 30 seconds of expiry
	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		return nil, err
	}

	// Create new JWT token
	var newToken Token
	err = newToken.Create(claims.Email, SECRET_KEY)
	if err != nil {
		return nil, err
	}

	res := &user.RefreshResponse{
		Token: newToken.Value,
	}

	return res, nil
}
