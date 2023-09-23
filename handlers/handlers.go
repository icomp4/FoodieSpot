package handlers

import (
	"foodSharer/controllers"
	"foodSharer/messages"
	"foodSharer/models"
	"github.com/gofiber/fiber/v2"
)

func HandleSignUp(c *fiber.Ctx) error {
	var user *models.User

	// Parsing body for login info
	if err := c.BodyParser(&user); err != nil {
		message := messages.ErrorMessage{
			Status:  "Failed",
			Message: "Failed to parse request body",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	if user.Username == "" || user.Password == "" {
		message := messages.ErrorMessage{
			Status:  "Failed",
			Message: "Fields must not be blank",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Added level of abstraction, see *\controllers\userController to see under the hood
	controllers.SignUp(user)
	message := messages.SuccessMessage{
		Status:  "Success",
		Message: "User " + user.Username + " has successfully been created",
		User:    user,
	}
	c.JSON(message)
	return c.SendStatus(200)
}
