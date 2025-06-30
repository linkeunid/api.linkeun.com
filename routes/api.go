package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"github.com/linkeunid/api.linkeun.com/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Prefix("api").Group(func(router route.Router) {
		router.Get("/users", userController.Index)
		router.Get("/users/{id}", userController.Show)
	})

	shortUrlController := controllers.NewShortUrlController()
	facades.Route().Group(func(router route.Router) {
		router.Get("/s", shortUrlController.Index)
		router.Get("/s/{id}", shortUrlController.Show)
		router.Post("/s", shortUrlController.Store)
		router.Put("/s/{id}", shortUrlController.Update)
		router.Delete("/s/{id}", shortUrlController.Destroy)
	})
}
