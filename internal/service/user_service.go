package service

import (
	"context"
	"errors"

	"github.com/getsentry/sentry-go"
	"github.com/linkeunid/api.linkeun.com/internal/models"
	"github.com/linkeunid/api.linkeun.com/internal/repository"
)

type UserService struct {
	repo      repository.UserRepository
	sentryHub *sentry.Hub
}

func NewUserService(repo repository.UserRepository, sentryHub *sentry.Hub) *UserService {
	return &UserService{
		repo:      repo,
		sentryHub: sentryHub,
	}
}

func (s UserService) CreateUser(ctx context.Context, user *models.CreateUserRequest) error {
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

func (s UserService) GetAll(ctx context.Context, limit, offset int) (*repository.UserListData, error) {
	users, err := s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	return users, nil
}

func (s UserService) GetByID(ctx context.Context, id uint64) (*models.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
func (s UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
func (s UserService) Update(ctx context.Context, id uint64, user *models.UpdateUserRequest) error {
	userFound, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return err
	}

	if userFound == nil {
		return errors.New("user not found")
	}

	if user.Password != "" {
		userFound.Password = user.Password
	}

	updateUserData := &models.User{
		ID:        userFound.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Password:  userFound.Password,
		CreatedAt: userFound.CreatedAt,
	}

	if err := s.repo.Update(ctx, updateUserData); err != nil {
		s.sentryHub.CaptureException(err)
		return err
	}

	return nil
}

func (s UserService) Delete(ctx context.Context, id uint64) error {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	return s.repo.Delete(ctx, id)
}
