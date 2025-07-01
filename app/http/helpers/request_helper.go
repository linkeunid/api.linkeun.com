package helpers

import "github.com/goravel/framework/contracts/http"

func AbortWithJsonResponse(ctx http.Context, code int, message string, data any, err any) {
	ctx.Request().AbortWithStatusJson(code, http.Json{
		"code":    code,
		"message": message,
		"data":    data,
		"error":   err,
	})
}

func AbortWithStatusResponse(ctx http.Context, code int) {
	ctx.Request().AbortWithStatus(code)
}
