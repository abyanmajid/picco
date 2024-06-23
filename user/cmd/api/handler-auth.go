package main

import (
	"context"
	"time"

	user "github.com/abyanmajid/codemore.io/user/proto"
	"github.com/abyanmajid/codemore.io/user/utils"
	"github.com/golang-jwt/jwt/v4"
)

func (api *Config) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	requestPayload := LoginPayload{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	// Check if email exists
	u, err := api.DB.GetUserByEmail(ctx, requestPayload.Email)
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
	err = token.Create(u.Email, api.SecretKey)
	if err != nil {
		return nil, err
	}

	res := &user.LoginResponse{
		Token:          token.Value,
		ExpirationTime: token.ExpiresAt.Unix(),
	}

	return res, nil
}

func (api *Config) Refresh(ctx context.Context, req *user.RefreshRequest) (*user.RefreshResponse, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		req.GetRefreshToken(),
		claims,
		func(token *jwt.Token) (any, error) {
			return api.SecretKey, nil
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
	err = newToken.Create(claims.Email, api.SecretKey)
	if err != nil {
		return nil, err
	}

	res := &user.RefreshResponse{
		Token:          newToken.Value,
		ExpirationTime: newToken.ExpiresAt.Unix(),
	}

	return res, nil
}
