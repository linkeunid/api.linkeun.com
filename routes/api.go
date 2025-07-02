package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"github.com/linkeunid/api.linkeun.com/app/http/controllers"
	"github.com/linkeunid/api.linkeun.com/app/http/middleware"
)

func Api() {
	facades.Route().Prefix("api").Group(func(router route.Router) {
		userController := controllers.NewUserController()
		router.Prefix("users").Middleware(middleware.Jwt(nil)).Group(func(router route.Router) {
			router.Get("/", userController.Index)
			router.Get("/{id}", userController.Show)
			router.Post("/", userController.Store)
			router.Patch("/{id}", userController.Update)
			router.Delete("/{id}", userController.Destroy)
			router.Get("/profile", userController.Profile)
		})

		authController := controllers.NewAuhController()
		router.Prefix("auth").Group(func(router route.Router) {
			router.Post("/login", authController.Login)
			router.Post("/register", authController.Register)
			router.Post("/logout", authController.Logout)
			router.Get("/verify/{token}", authController.Verify)
		})

		shortUrlController := controllers.NewShortUrlController()
		shortUrlControllerExceptPaths := middleware.ExceptPaths{
			"/api/s/": []string{"POST"},
		}
		router.Get("/s/{shortCode}", shortUrlController.Redirect)
		router.Prefix("s").Middleware(middleware.Jwt(&shortUrlControllerExceptPaths)).Group(func(router route.Router) {
			router.Post("/", shortUrlController.Store)
			router.Get("/", shortUrlController.Index)
			router.Get("/{shortCode}/detail", shortUrlController.Show)
			router.Put("/{shortCode}", shortUrlController.Update)
			router.Delete("/{shortCode}", shortUrlController.Destroy)
		})

	})
}
