package service

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type ToolService struct {
	logger *slog.Logger
}

func NewToolService(logger *slog.Logger) *ToolService {
	return &ToolService{
		logger: logger,
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
		return nil, err
	}

	var ipInfo map[string]interface{}
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return nil, err
	}

	return &ipInfo, nil
}
