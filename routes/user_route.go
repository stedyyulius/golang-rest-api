package routes

import (
	"golang-rest-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/user/:userId", controllers.GetAUser)
	app.Post("/user", controllers.CreateUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}
