package routes

import (
	"jwt/controllers"

	"github.com/gofiber/fiber/v2"
)

// Setup Routes
func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}