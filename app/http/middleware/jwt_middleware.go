package middleware

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
)

func Jwt() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			helpers.AbortWithJsonResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, "missing token")
			return
		}

		// The token normally comes in format `Bearer <token>`, so we need to remove `Bearer ` prefix.
		token = strings.Replace(token, "Bearer ", "", 1)

		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			helpers.AbortWithJsonResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, err.Error())
			return
		}

		ctx.Request().Next()
	}
}
