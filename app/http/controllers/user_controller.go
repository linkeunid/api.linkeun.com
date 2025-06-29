package controllers

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
	"github.com/linkeunid/api.linkeun.com/app/http/resources"
	"github.com/linkeunid/api.linkeun.com/app/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (r *UserController) Index(ctx http.Context) http.Response {
	paginateInput := helpers.ParsePaginationRequest(ctx)

	users, total, err := r.userService.GetUsers(ctx, paginateInput.Page, paginateInput.PerPage, paginateInput.Sort, paginateInput.Search, paginateInput.Filters)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	}

	data := resources.MakeUserCollection(users)
	meta := helpers.MakePaginationMeta(total, paginateInput.Page, paginateInput.PerPage)

	return helpers.JsonPaginateResponse(ctx, http.StatusOK, "success", data, meta, nil)
}

func (r *UserController) Show(ctx http.Context) http.Response {
	userId := ctx.Request().RouteInt("id")

	user, err := r.userService.GetUser(ctx, strconv.Itoa(userId))
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusNotFound, "user not found", nil, nil)
	}

	data := resources.MakeUserResource(user).ToArray()

	return helpers.JsonResponse(ctx, http.StatusOK, "success", data, nil)
}
