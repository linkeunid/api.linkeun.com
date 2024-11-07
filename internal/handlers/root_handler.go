package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	"github.com/linkeunid/api.linkeun.com/pkg/utils"
)

type RootHandler struct {
	Logger *slog.Logger
}

func NewRootHandler(logger *slog.Logger) *RootHandler {
	return &RootHandler{
		Logger: logger,
	}
}

// Register
func (h *RootHandler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.HandleHome)           // GET /
	r.Get("/error", h.HandleTestError) // GET /error
	r.Get("/panic", h.HandleTestPanic) // GET /panic
}

func (h *RootHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusOK,
		Data:    "Hello World!!!",
		Error:   false,
		Message: "success",
	})
}

func (h *RootHandler) HandleTestError(w http.ResponseWriter, r *http.Request) {
	hub := sentry.GetHubFromContext(r.Context())
	hub.CaptureException(errors.New("test error"))

	utils.WriteJSONResponse(w, &utils.ResponseOpts{
		Code:    http.StatusInternalServerError,
		Data:    nil,
		Error:   true,
		Message: "error",
	})
}

func (h *RootHandler) HandleTestPanic(w http.ResponseWriter, r *http.Request) {
	panic("test panic")
}
