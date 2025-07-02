package middleware

import (
	"slices"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
)

type ExceptPaths map[string][]string

func Jwt(exceptPaths *ExceptPaths) http.Middleware {
	return func(ctx http.Context) {
		isExcepted := false
		if exceptPaths != nil {
			if methods, ok := (*exceptPaths)[ctx.Request().Path()]; ok {
				if slices.Contains(methods, ctx.Request().Method()) {
					isExcepted = true
				}
			}
		}

		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			if isExcepted {
				ctx.Request().Next()
				return
			}

			helpers.AbortWithJsonResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, "missing token")
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)

		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			helpers.AbortWithJsonResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, err.Error())
			return
		}

		ctx.Request().Next()
	}
}
