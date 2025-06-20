package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/huyvu/go-fiber-todolist/models"
	"github.com/huyvu/go-fiber-todolist/services"
	"github.com/huyvu/go-fiber-todolist/utils"
)

func GetUsers(c *fiber.Ctx) error {
	users := services.GetAllUsers()
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	users := services.GetAllUsers()
	// Tại sao dùng &user => Vì BodyParser cần thay đổi giá trị của struct
	// if err := c.BodyParser(&user); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	// } => đây là cách viết ngắn gọn của đoạn code bên dưới
	var err error
	err = c.BodyParser(&user) // => Gọi hàm parse JSON và gán lỗi (nếu có) vào biến
	// fmt.Println("err:", err)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	for _, u := range users {
		if u.Username == user.Username {
			return c.Status(400).JSON(fiber.Map{"error": "User already exists"})
		}
	}

	newUser := services.CreateUser(user)
	return c.Status(201).JSON(newUser)
}

func Login(c *fiber.Ctx) error {
	var input models.User
	users := services.GetAllUsers()
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	for _, u := range users {
		if u.Username == input.Username && u.Password == input.Password {
			token, err := utils.GenerateJWT(u.Username)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
			}
			return c.JSON(fiber.Map{"token": token, "user": input})
		}
	}

	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}
