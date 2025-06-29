package helpers

import "github.com/goravel/framework/contracts/http"

func JsonResponse(ctx http.Context, code int, message string, data any, err any) http.Response {
	return ctx.Response().Status(code).Json(map[string]any{
		"code":    code,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func JsonPaginateResponse(ctx http.Context, code int, message string, data any, pagination map[string]any, err any) http.Response {
	return ctx.Response().Status(code).Json(map[string]any{
		"code":    code,
		"message": message,
		"data":    data,
		"meta":    pagination,
		"error":   err,
	})
}
