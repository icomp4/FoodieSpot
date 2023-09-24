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
func HandleLogin(c *fiber.Ctx) error {
	var user *models.User
	// Parsing body for login info
	if err := c.BodyParser(&user); err != nil {
		message := messages.ErrorMessage{
			Status:  "Login Failed",
			Message: "Failed to parse request body",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Attempting to log in via the controllers.Login method, returns a user and an error
	attemptLogin, err := controllers.Login(user.Username, user.Password)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Login Failed",
			Message: "Account details do not match",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// If no error is returned, login has been successful
	message := messages.SuccessMessage{
		Status:  "Login successful",
		Message: "User " + attemptLogin.Username + " has successfully logged in !",
		User:    attemptLogin,
	}
	c.JSON(message)
	return c.SendStatus(200)

}
