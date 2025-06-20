package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/huyvu/go-fiber-todolist/middleware"
	"github.com/huyvu/go-fiber-todolist/routes"
	"github.com/huyvu/go-fiber-todolist/storage"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	// Lấy biến môi trường
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // fallback nếu PORT không được set
	}
	// Check variable is pointers
	// fmt.Println("is pointer:", reflect.TypeOf(app).Kind() == reflect.Ptr)
	storage.LoadData()

	routes.AuthRoutes(app)
	app.Use(middleware.JWTProtected())
	// Mục đích là đăng ký (mount) các route / endpoint cho module user.
	routes.UserRoutes(app)

	app.Listen(":" + port)
}
