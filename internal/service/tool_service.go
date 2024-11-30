package service

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
)

type ToolService struct {
	logger    *slog.Logger
	sentryHub *sentry.Hub
}

func NewToolService(logger *slog.Logger, sentryHub *sentry.Hub) *ToolService {
	return &ToolService{
		logger:    logger,
		sentryHub: sentryHub,
	}
}

func (s ToolService) GetIPInfo(ctx context.Context, clientIP string) (*map[string]interface{}, error) {
	resp, err := http.Get("https://ipinfo.info/ip_api.php?ip=" + clientIP)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	var ipInfo map[string]interface{}
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		s.sentryHub.CaptureException(err)
		return nil, err
	}

	return &ipInfo, nil
}
