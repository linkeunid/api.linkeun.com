package controllers

import (
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
	"github.com/linkeunid/api.linkeun.com/app/http/requests"
	"github.com/linkeunid/api.linkeun.com/app/http/resources"
	"github.com/linkeunid/api.linkeun.com/app/services"
)

type AuhController struct {
	authService *services.AuthService
}

func NewAuhController() *AuhController {
	return &AuhController{
		authService: services.NewAuthService(),
	}
}

func (r *AuhController) Login(ctx http.Context) http.Response {
	var loginRequest requests.LoginRequest
	errors, _ := ctx.Request().ValidateRequest(&loginRequest)
	if errors != nil {
		return helpers.JsonResponse(ctx, http.StatusUnprocessableEntity, "error", nil, errors.All())
	}

	user, token, err := r.authService.Login(ctx, &loginRequest)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil, err)
	}

	response := map[string]interface{}{
		"user":       resources.MakeUserResource(user).ToArray(),
		"token":      *token,
		"expires_in": facades.Config().Get("auth.guards.user.ttl", 24*time.Hour).(time.Duration).Seconds(),
	}

	return helpers.JsonResponse(ctx, http.StatusOK, "success", response, nil)
}

func (r *AuhController) Register(ctx http.Context) http.Response {
	var registerRequest requests.RegisterRequest
	errors, _ := ctx.Request().ValidateRequest(&registerRequest)
	if errors != nil {
		return helpers.JsonResponse(ctx, http.StatusUnprocessableEntity, "error", nil, errors.All())
	}

	user, err := r.authService.Register(ctx, &registerRequest)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	}

	return helpers.JsonResponse(ctx, http.StatusOK, "please check your email to verify your account", resources.MakeUserResource(user).ToArray(), nil)
}

func (r *AuhController) Logout(ctx http.Context) http.Response {
	return nil
}

func (r *AuhController) Verify(ctx http.Context) http.Response {
	token := ctx.Request().Route("token")

	user, authToken, err := r.authService.VerifyEmail(ctx, token)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusUnauthorized, err.Error(), nil, nil)
	}

	response := map[string]interface{}{
		"user":       resources.MakeUserResource(user).ToArray(),
		"token":      *authToken,
		"expires_in": facades.Config().Get("auth.guards.user.ttl", 24*time.Hour).(time.Duration).Seconds(),
	}

	return helpers.JsonResponse(ctx, http.StatusOK, "success", response, nil)
}
