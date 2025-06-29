package services

import (
	"context"
	"strings"

	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers(ctx context.Context, page int, perPage int, sortBy string, search string, filters map[string]string) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	query := facades.Orm().WithContext(ctx).Query().Model(&models.User{})

	if search != "" {
		searchTerm := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", searchTerm, searchTerm)
	}

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if sortBy != "" {
		query = query.Order(sortBy)
	}

	err := query.Paginate(page, perPage, &users, &total)

	return users, total, err
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := facades.Orm().WithContext(ctx).Query().FindOrFail(&user, id)

	return &user, err
}
