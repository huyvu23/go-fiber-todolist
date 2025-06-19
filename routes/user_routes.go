package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/handlers"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Get("/", handlers.GetUsers)
	user.Post("/", handlers.CreateUser)
}
