package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/routes"
	"github.com/huyvu/go-fiber-todolist/storage"
)

func main() {
	app := fiber.New()
	storage.LoadData()
	routes.RegisterUserRoutes(app)
	app.Listen(":3000")
}
