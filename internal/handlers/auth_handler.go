package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/linkeunid/api.linkeun.com/internal/models"
	"github.com/linkeunid/api.linkeun.com/internal/service"
	"github.com/linkeunid/api.linkeun.com/pkg/utils"
)

type AuthHandler struct {
	service   *service.AuthService
	validator *validator.Validate
	Logger    *slog.Logger
}

func NewAuthHandler(service *service.AuthService, logger *slog.Logger) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator.New(),
		Logger:    logger,
	}
}

// Register
func (h *AuthHandler) RegisterRoutesV1(r chi.Router) {
	r.Post("/register", h.Register) // POST /v1/register
	r.Post("/signin", h.SignIn)     // POST /v1/signin
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
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
	err := h.service.Register(ctx, &user)
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
		Data:    "register successfully",
		Error:   false,
		Message: "success",
	})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var signReq models.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&signReq); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "invalid input",
		})
		return
	}

	if err := h.validator.Struct(signReq); err != nil {
		utils.WriteJSONResponse(w, &utils.ResponseOpts{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Error:   true,
			Message: "validation failed: " + err.Error(),
		})
		return
	}

	ctx := r.Context()
	signData, err := h.service.SignIn(ctx, &signReq)
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
		Data:    signData,
		Error:   false,
		Message: "success",
	})
}
