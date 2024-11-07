package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/linkeunid/api.linkeun.com/internal/service"
	"github.com/linkeunid/api.linkeun.com/pkg/utils"
)

type ToolHandler struct {
	logger  *slog.Logger
	service *service.ToolService
}

func NewToolHandler(logger *slog.Logger, service *service.ToolService) *ToolHandler {
	return &ToolHandler{
		service: service,
		logger:  logger,
	}
}

// Register
func (h *ToolHandler) RegisterRoutes(r chi.Router) {
	r.Get("/ip", h.GetIPInfo) // GET /ip
}

func (h *ToolHandler) GetIPInfo(w http.ResponseWriter, r *http.Request) {
	clientIP := r.Header.Get("X-FORWARDED-FOR")
	if clientIP == "" {
		clientIP = "127.0.0.1"
	}

	ctx := r.Context()
	ipInfo, err := h.service.GetIPInfo(ctx, clientIP)
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
		Data:    ipInfo,
		Error:   false,
		Message: "success",
	})
}
