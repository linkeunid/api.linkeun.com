package controllers

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
	"github.com/linkeunid/api.linkeun.com/app/http/requests"
	"github.com/linkeunid/api.linkeun.com/app/http/resources"
	"github.com/linkeunid/api.linkeun.com/app/models"
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
	return helpers.JsonResponse(ctx, http.StatusNotImplemented, "not implemented", nil, nil)
	// paginateInput := helpers.ParsePaginationRequest(ctx)

	// users, total, err := r.userService.GetUsers(ctx, paginateInput.Page, paginateInput.PerPage, paginateInput.Sort, paginateInput.Search, paginateInput.Filters)
	// if err != nil {
	// 	return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	// }

	// data := resources.MakeUserCollection(users)
	// meta := helpers.MakePaginationMeta(total, paginateInput.Page, paginateInput.PerPage)

	// return helpers.JsonPaginateResponse(ctx, http.StatusOK, "success", data, meta, nil)
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return helpers.JsonResponse(ctx, http.StatusNotImplemented, "not implemented", nil, nil)
	// userId := ctx.Request().RouteInt("id")

	// user, err := r.userService.GetUser(ctx, strconv.Itoa(userId))
	// if err != nil {
	// 	return helpers.JsonResponse(ctx, http.StatusNotFound, "user not found", nil, nil)
	// }

	// data := resources.MakeUserResource(user).ToArray()

	// return helpers.JsonResponse(ctx, http.StatusOK, "success", data, nil)
}

func (r *UserController) Store(ctx http.Context) http.Response {
	return helpers.JsonResponse(ctx, http.StatusNotImplemented, "not implemented", nil, nil)
	// var request requests.RegisterRequest
	// if err := ctx.Request().Bind(&request); err != nil {
	// 	return helpers.JsonResponse(ctx, http.StatusUnprocessableEntity, "error", nil, err)
	// }

	// hashedPassword, err := facades.Hash().Make(request.Password)
	// if err != nil {
	// 	return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	// }

	// user, err := r.userService.CreateUser(ctx.Context(), &models.User{
	// 	Name:     request.Name,
	// 	Email:    request.Email,
	// 	Password: hashedPassword,
	// })

	// if err != nil {
	// 	return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	// }

	// return helpers.JsonResponse(ctx, http.StatusCreated, "success", user, nil)
}

func (r *UserController) Update(ctx http.Context) http.Response {
	userId := ctx.Request().RouteInt("id")
	var request requests.UpdateUserRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return helpers.JsonResponse(ctx, http.StatusUnprocessableEntity, "error", nil, err)
	}

	user, err := r.userService.UpdateUser(ctx.Context(), strconv.Itoa(userId), &request)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	}

	return helpers.JsonResponse(ctx, http.StatusOK, "success", user, nil)
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Profile(ctx http.Context) http.Response {
	var user models.User
	if err := facades.Auth(ctx).User(&user); err != nil {
		return helpers.JsonResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, err)
	}

	data := resources.MakeUserResource(&user).ToArray()

	return helpers.JsonResponse(ctx, http.StatusOK, "success", data, nil)
}
