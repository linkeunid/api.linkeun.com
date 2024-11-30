package main

import (
	"context"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	"github.com/linkeunid/api.linkeun.com/internal/handlers"
	"github.com/linkeunid/api.linkeun.com/internal/service"
)

func (app *App) setupRoutes(r chi.Router) {
	ctx := sentry.SetHubOnContext(context.Background(), sentry.CurrentHub().Clone())
	hub := sentry.GetHubFromContext(ctx)

	rootHandler := handlers.NewRootHandler(app.logger)
	rootHandler.RegisterRoutes(r) // Register root routes

	// userRepo := repository.NewUserRepository(app.logger, app.db, app.bcrypt)

	r.Route("/v1", func(v1 chi.Router) {
		// v1.Route("/auth", func(authRouter chi.Router) {
		// 	authService := service.NewAuthService(app.logger, userRepo, hub, app.bcrypt)
		// 	authHandler := handlers.NewAuthHandler(app.logger, authService)
		// 	authHandler.RegisterRoutesV1(authRouter)
		// })

		// v1.Route("/users", func(userRouter chi.Router) {
		// 	userRouter.Use(middlewares.AuthMiddleware)

		// 	userService := service.NewUserService(app.logger, userRepo, hub) // Initialize the user service
		// 	userHandler := handlers.NewUserHandler(app.logger, userService)  // Initialize the handler
		// 	userHandler.RegisterRoutesV1(userRouter)                         // Register the user routes
		// })

		v1.Route("/tools", func(toolsRouter chi.Router) {
			toolService := service.NewToolService(app.logger, hub)
			toolHandler := handlers.NewToolHandler(app.logger, toolService)
			toolHandler.RegisterRoutes(toolsRouter)
		})
	})
}
