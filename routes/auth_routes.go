package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/handlers"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", handlers.CreateUser)
	auth.Post("/login", handlers.Login)
}
