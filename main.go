package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/routes"
	"github.com/huyvu/go-fiber-todolist/storage"
)

func main() {
	app := fiber.New()
	// Check variable is pointers
	// fmt.Println("is pointer:", reflect.TypeOf(app).Kind() == reflect.Ptr)
	storage.LoadData()
	// Mục đích là đăng ký (mount) các route / endpoint cho module user.
	routes.RegisterUserRoutes(app)
	app.Listen(":3000")
}
