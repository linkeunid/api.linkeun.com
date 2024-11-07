package service

import (
	"context"
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/golang-jwt/jwt/v5"
	"github.com/linkeunid/api.linkeun.com/internal/models"
	"github.com/linkeunid/api.linkeun.com/internal/repository"
	"github.com/linkeunid/api.linkeun.com/pkg/env"
)

type AuthService struct {
	repo      repository.UserRepository
	sentryHub *sentry.Hub
}

func NewAuthService(repo repository.UserRepository, sentryHub *sentry.Hub) *AuthService {
	return &AuthService{
		repo:      repo,
		sentryHub: sentryHub,
	}
}

func (s AuthService) Register(ctx context.Context, user *models.CreateUserRequest) error {
	createUserData := &models.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := s.repo.Create(ctx, createUserData); err != nil {
		s.sentryHub.CaptureException(err)
		return err
	}

	return nil
}

func (s AuthService) SignIn(ctx context.Context, user *models.SignInRequest) (*models.SignInResponse, error) {
	userFound, err := s.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	if userFound == nil {
		return nil, errors.New("user not found")
	}

	expTime, err := time.ParseDuration(env.GetString("JWT_EXPIRES", "1") + "h")
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"sub":        userFound.ID,
		"user":       userFound,
		"exp":        time.Now().Add(expTime * 1).Unix(),
		"iat":        time.Now().Unix(),
		"iss":        "api.linkeun.com",
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(env.GetString("JWT_SECRET", "")))
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	return &models.SignInResponse{
		User:        userFound,
		AccessToken: tokenString,
		ExpiresIn:   int64(expTime.Seconds()),
	}, nil
}