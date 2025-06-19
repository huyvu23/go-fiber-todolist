package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/models"
	"github.com/huyvu/go-fiber-todolist/services"
)

func GetUsers(c *fiber.Ctx) error {
	users := services.GetAllUsers()
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	// Tại sao dùng &user => Vì BodyParser cần thay đổi giá trị của struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	newUser := services.CreateUser(user)
	return c.Status(201).JSON(newUser)
}
