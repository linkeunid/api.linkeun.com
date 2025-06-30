package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateUserRequest struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *UpdateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateUserRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "sometimes|required",
		"email":    "sometimes|required|email|unique:users,email," + ctx.Request().Route("id"),
		"password": "sometimes|required|min:6",
	}
}

func (r *UpdateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
