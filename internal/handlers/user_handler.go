package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/linkeunid/api.linkeun.com/internal/models"
	"github.com/linkeunid/api.linkeun.com/internal/service"
	"github.com/linkeunid/api.linkeun.com/pkg/utils"
)

type UserHandler struct {
	service   *service.UserService
	validator *validator.Validate
	Logger    *slog.Logger
}

func NewUserHandler(service *service.UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service:   service,
		validator: validator.New(),
		Logger:    logger,
	}
}

// Register
func (h *UserHandler) RegisterRoutesV1(r chi.Router) {
	r.Post("/", h.CreateUser)     // POST /v1/users
	r.Get("/", h.GetAllUser)      // GET /v1/users
	r.Get("/{id}", h.GetUserByID) // GET /v1/users/{id}
	// r.Get("/users/{email}", h.GetUserByEmail) // GET /v1/users/{id}
	r.Patch("/{id}", h.UpdateUser)  // PUT /v1/users/{id}
	r.Delete("/{id}", h.DeleteUser) // DELETE /v1/users/{id}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "invalid input",
		})
		return
	}

	if err := h.validator.Struct(user); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "validation failed: " + err.Error(),
		})
		return
	}

	ctx := r.Context()
	err := h.service.CreateUser(ctx, &user)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusCreated,
		Data:    "User created successfully",
		Error:   false,
		Message: "success",
	})
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	ctx := r.Context()
	data, err := h.service.GetAll(ctx, limit, offset)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	if len(data.Users) == 0 {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusOK,
			Data:    nil,
			Error:   false,
			Message: "data not found",
		})
		return
	}

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:        http.StatusOK,
		Data:        data.Users,
		Error:       false,
		Message:     "success",
		Limit:       &limit,
		Offset:      &offset,
		Total:       &data.TotalCount,
		TotalPages:  &data.TotalPages,
		CurrentPage: &data.CurrentPage,
		HasNext:     &data.HasNext,
		HasPrev:     &data.HasPrev,
	})
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "invalid id",
		})
		return
	}

	ctx := r.Context()
	user, err := h.service.GetByID(ctx, id)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusNotFound,
			Data:    nil,
			Error:   true,
			Message: "user not found",
		})
		return
	}

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusOK,
		Data:    user,
		Error:   false,
		Message: "success",
	})
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	ctx := r.Context()
	user, err := h.service.GetByEmail(ctx, email)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusNotFound,
			Data:    nil,
			Error:   true,
			Message: "user not found",
		})
		return
	}

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusOK,
		Data:    user,
		Error:   false,
		Message: "success",
	})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "invalid input",
		})
		return
	}

	if err := h.validator.Struct(user); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "validation failed: " + err.Error(),
		})
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "invalid id",
		})
		return
	}

	ctx := r.Context()
	err = h.service.Update(ctx, id, &user)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusOK,
		Data:    "User update successfully",
		Error:   false,
		Message: "success",
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "invalid id",
		})
		return
	}

	ctx := r.Context()
	if err := h.service.Delete(ctx, id); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Error:   true,
			Message: err.Error(),
		})
		return
	}

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusOK,
		Data:    "User deleted successfully",
		Error:   false,
		Message: "success",
	})
}
