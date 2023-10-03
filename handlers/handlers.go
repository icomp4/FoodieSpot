package handlers

import (
	"fmt"
	"foodSharer/controllers"
	"foodSharer/messages"
	"foodSharer/models"
	"foodSharer/session"
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

	// getting the current session store
	sess, err := session.Store.Get(c)
	if err != nil {
		panic(err)
	}

	// attempts to set the session's USERID, USERNAME, and AUTHORIZED
	sessionFailed := session.SetSession(sess, attemptLogin)
	if sessionFailed != nil {
		message := messages.ErrorMessage{
			Status:  "Login Failed",
			Message: "Failed to save session",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	// User successfully logged in, return user and status
	message := messages.SuccessMessage{
		Status:  "Login successful",
		Message: "User " + attemptLogin.Username + " has successfully logged in !",
		User:    attemptLogin,
	}
	c.JSON(message)
	return c.SendStatus(200)
}
func HandleLogout(c *fiber.Ctx) error {
	// Getting current session info
	sess, err := session.Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Logout Failed",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	// Storing the current user's name
	username := sess.Get("Username").(string)

	// Deleting the cookie, (logging out)
	destroy := sess.Destroy()
	if err != nil {
		panic(destroy)
	}
	message := messages.SuccessMessage{
		Status:  "Logout Successful",
		Message: "User " + username + " has successfully logged out",
	}
	c.JSON(message)
	return c.SendStatus(400)
}
func HandleGetLocation(c *fiber.Ctx) error {

	// Attempting to get current session information
	sess, err := session.Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Failed to get ip",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	//Checking if the current user is authorized (logged in)
	if sess.Get("Authorized") == false {
		message := messages.ErrorMessage{
			Status:  "Failed to get ip",
			Message: "User not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Retrieving the ip address, and passing it to GetLocation func
	ip := c.IP()
	fmt.Println(ip)
	response, err := controllers.GetLocation(ip)
	if err != nil || response.Success == false {
		message := messages.ErrorMessage{
			Status:  "Failed to get ip",
			Message: "Invalid address",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Formatting the proper response for a successful request
	message := messages.LocationMessage{
		Status:   "Success",
		Message:  "Successfully retrieved geodata",
		Location: response,
	}
	c.JSON(message)
	return c.SendStatus(200)

}
