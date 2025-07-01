package services

import (
	"errors"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
	"github.com/linkeunid/api.linkeun.com/app/http/requests"
	"github.com/linkeunid/api.linkeun.com/app/mails"
	"github.com/linkeunid/api.linkeun.com/app/models"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		userService: NewUserService(),
	}
}

func (s *AuthService) Login(ctx http.Context, request *requests.LoginRequest) (*models.User, *string, error) {
	loginError := errors.New("username or password is incorrect")
	user, err := s.userService.GetUserByUsername(ctx.Context(), request.Username)
	if err != nil {
		return nil, nil, loginError
	}

	if !user.IsVerified {
		return nil, nil, errors.New("email not verified")
	}

	if !facades.Hash().Check(request.Password, user.Password) {
		return nil, nil, loginError
	}

	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return nil, nil, err
	}

	return user, &token, nil
}

func (s *AuthService) Register(ctx http.Context, request *requests.RegisterRequest) (*models.User, error) {
	hashedPassword, err := facades.Hash().Make(request.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.CreateUser(ctx.Context(), &models.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateRandomString(32)
	if err != nil {
		return nil, err
	}

	verification := models.EmailVerification{
		UserId:    uint64(user.ID),
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := facades.Orm().Query().Create(&verification); err != nil {
		return nil, err
	}

	err = facades.Mail().To([]string{user.Email}).Send(mails.NewVerifyEmail(user, token))
	if err != nil {
		facades.Log().Error("failed to send verification email: ", err.Error())
	}

	return user, nil
}

func (s *AuthService) VerifyEmail(ctx http.Context, token string) (*models.User, *string, error) {
	var verification models.EmailVerification
	err := facades.Orm().Query().Where("token", token).FirstOrFail(&verification)
	if err != nil {
		return nil, nil, errors.New("invalid verification token")
	}

	if time.Now().After(verification.ExpiresAt) {
		return nil, nil, errors.New("verification token expired")
	}

	var user models.User
	err = facades.Orm().Query().FindOrFail(&user, verification.UserId)
	if err != nil {
		return nil, nil, errors.New("user not found")
	}

	user.IsVerified = true
	if err := facades.Orm().Query().Save(&user); err != nil {
		return nil, nil, err
	}

	if _, err := facades.Orm().Query().Delete(&verification); err != nil {
		// Log the error but don't fail the request, as the user is already verified
		facades.Log().Errorf("failed to delete verification token: %v", err)
	}

	authToken, err := facades.Auth(ctx).Login(&user)
	if err != nil {
		return nil, nil, err
	}

	return &user, &authToken, nil
}

func (s *AuthService) Logout(ctx http.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func (s *AuthService) Verify(ctx http.Context, user *models.User) (*models.User, error) {
	return nil, nil
}
