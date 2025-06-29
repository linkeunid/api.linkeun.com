package routes

import (
	"github.com/goravel/framework/facades"

	"github.com/linkeunid/api.linkeun.com/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users", userController.Index)
	facades.Route().Get("/users/{id}", userController.Show)
}
