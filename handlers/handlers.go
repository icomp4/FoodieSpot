package handlers

import (
	"foodSharer/controllers"
	"foodSharer/models"
	"github.com/gofiber/fiber/v2"
)

func HandleSignUp(c *fiber.Ctx) error {
	var user *models.User

	// Parsing body for login info
	if err := c.BodyParser(&user); err != nil {
		c.SendString("Error parsing body")
		return c.SendStatus(404)
	}
	if user.Username == "" || user.Password == "" {
		c.SendString("Fields must not be blank")
		return c.SendStatus(404)
	}

	// Added level of abstraction, see *\controllers\userController to see under the hood
	controllers.SignUp(user)
	c.SendString("User " + user.Username + " successfully created !")
	return c.SendStatus(201)
}
