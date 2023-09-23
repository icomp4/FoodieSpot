package handlers

import (
	"foodSharer/controllers"
	"foodSharer/models"
	"github.com/gofiber/fiber/v2"
)

func HandleSignUp(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		c.SendString("Error parsing body")
		return c.SendStatus(404)
	}
	if user.Username == "" || user.Password == "" {
		c.SendString("Fields must not be blank")
		return c.SendStatus(404)
	}
	controllers.SignUp(user)
	c.SendString("User " + user.Username + " successfully created !")
	return c.SendStatus(201)
}
