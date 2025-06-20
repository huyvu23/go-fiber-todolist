package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/handlers"
)

// => Tại sao lại viết *fiber.App => bởi vì khi khởi tạo fiber trả về 1 con trỏ và truyền xuống RegisterUserRoutes
func UserRoutes(app *fiber.App) {

	// Tạo một nhóm các route (route group) bắt đầu bằng /user
	user := app.Group("/user")

	// user.Get("/", func(c *fiber.Ctx) error {
	// 	return handlers.GetUsers(c)
	// }) => Đây chính là cách viết tường minh
	user.Get("/", handlers.GetUsers)
}
