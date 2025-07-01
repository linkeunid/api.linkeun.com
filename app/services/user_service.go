package services

import (
	"context"
	"strings"

	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/requests"
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

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := facades.Orm().WithContext(ctx).Query().Where("username = ?", username).FirstOrFail(&user)

	return &user, err
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := facades.Orm().WithContext(ctx).Query().Where("email", email).FirstOrFail(&user)

	return &user, err
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := facades.Orm().WithContext(ctx).Query().Create(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, request *requests.UpdateUserRequest) (*models.User, error) {
	var user models.User
	err := facades.Orm().WithContext(ctx).Query().FindOrFail(&user, id)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		hashedPassword, err := facades.Hash().Make(request.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	err = facades.Orm().WithContext(ctx).Query().Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := facades.Orm().WithContext(ctx).Query().Model(&models.User{}).Where("id = ?", user.ID).Delete()

	return user, err
}
