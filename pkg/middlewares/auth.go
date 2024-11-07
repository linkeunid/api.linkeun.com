package middlewares

import (
	"net/http"
	"strings"

	"github.com/linkeunid/api.linkeun.com/pkg/env"
	"github.com/linkeunid/api.linkeun.com/pkg/jwt"
	"github.com/linkeunid/api.linkeun.com/pkg/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has an Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.WriteJSONResponse(w, &utils.ResponseOpts{
				Code:    http.StatusUnauthorized,
				Data:    nil,
				Error:   false,
				Message: "bearer token required",
			})
			return
		}

		// Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		_, err := jwt.ValidateJWT(tokenString, env.GetString("JWT_SECRET", ""))
		if err != nil {
			utils.WriteJSONResponse(w, &utils.ResponseOpts{
				Code:    http.StatusOK,
				Data:    nil,
				Error:   false,
				Message: "invalid token",
			})
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
