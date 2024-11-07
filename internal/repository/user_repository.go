package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/linkeunid/api.linkeun.com/internal/database"
	"github.com/linkeunid/api.linkeun.com/internal/models"
	"github.com/linkeunid/api.linkeun.com/pkg/bcrypt"
	"github.com/linkeunid/api.linkeun.com/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context, opts *utils.OrderingFilter) (*UserListData, error)
	GetByID(ctx context.Context, id uint64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, updatedUser *models.User) error
	Delete(ctx context.Context, id uint64) error
}

type UserListData struct {
	Users       []models.User
	TotalCount  int
	TotalPages  int
	CurrentPage int
	HasNext     bool
	HasPrev     bool
}

type UserRepositoryImpl struct {
	logger *slog.Logger
	db     *database.DB
	hash   *bcrypt.Bcrypt
}

func NewUserRepository(logger *slog.Logger, db *database.DB, hash *bcrypt.Bcrypt) UserRepository {
	// Run migrations
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
		os.Exit(1)
	}

	return &UserRepositoryImpl{
		logger: logger,
		db:     db,
		hash:   hash,
	}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	hashedPassword, err := repo.hash.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = repo.db.WithContext(ctx).Create(user).Error
	if err != nil {
		repo.logger.Error("failed to create user", "error", err)
	}

	return err
}

func (repo *UserRepositoryImpl) GetAll(ctx context.Context, opts *utils.OrderingFilter) (*UserListData, error) {
	var users []models.User

	opts.OrderBy = strings.TrimSpace(strings.ToLower(opts.OrderBy))
	opts.Order = strings.TrimSpace(strings.ToLower(opts.Order))

	if opts.OrderBy == "" {
		opts.OrderBy = "id"
	}

	if opts.Order == "" {
		opts.Order = "desc"
	}

	ordering := fmt.Sprintf("%s %s", opts.OrderBy, opts.Order)

	if err := repo.db.WithContext(ctx).Order(ordering).Limit(opts.Limit).Offset(opts.Offset).Find(&users).Error; err != nil {
		repo.logger.Error("failed to get users", "error", err)
		return nil, err
	}

	var totalCount int64
	if err := repo.db.WithContext(ctx).Model(&models.User{}).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := int((totalCount + int64(opts.Limit) - 1) / int64(opts.Limit))
	currentPage := opts.Offset/opts.Limit + 1
	hasNext := currentPage < totalPages
	hasPrev := currentPage > 1

	return &UserListData{
		Users:       users,
		TotalCount:  int(totalCount),
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		HasNext:     hasNext,
		HasPrev:     hasPrev,
	}, nil
}

func (repo *UserRepositoryImpl) GetByID(ctx context.Context, id uint64) (*models.User, error) {
	var user models.User
	if err := repo.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		repo.logger.Error("failed to get user", "error", err)

		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		repo.logger.Error("failed to get user by email", "error", err, "email", email)
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, updatedUser *models.User) error {
	if updatedUser.Password != "" {
		hashedPassword, err := repo.hash.HashPassword(updatedUser.Password)
		if err != nil {
			repo.logger.Error("failed to hash password", "error", err)
			return err
		}

		updatedUser.Password = hashedPassword
	}

	if err := repo.db.WithContext(ctx).Save(updatedUser).Error; err != nil {
		repo.logger.Error("failed to update user", "error", err)
		return err
	}

	return nil
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	if err := repo.db.WithContext(ctx).Delete(&models.User{}, id).Error; err != nil {
		repo.logger.Error("failed to delete user", "error", err)
		return err
	}

	return nil
}
