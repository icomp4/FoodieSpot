package handlers

import (
	"fmt"
	"foodSharer/controllers"
	"foodSharer/messages"
	"foodSharer/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func init() {
	store := session.New(session.Config{})
	Store = store
}
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
			Status:  "Error",
			Message: "Failed to parse request body",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Attempting to log in via the controllers.Login method, returns a user and an error
	attemptLogin, err := controllers.Login(user.Username, user.Password)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Account details do not match",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Getting the current session store
	sess, err := Store.Get(c)
	if err != nil {
		panic(err)
	}

	// User successfully logged in, return user and status
	message := messages.SuccessMessage{
		Status:  "Success",
		Message: "User " + attemptLogin.Username + " has successfully logged in !",
		User:    attemptLogin,
	}
	// Attempts to set the session's USERID, USERNAME, and AUTHORIZED
	idAsString := strconv.Itoa(int(attemptLogin.ID))
	sess.Set("UserID", idAsString)
	sess.Set("Username", attemptLogin.Username)
	sess.Set("Authorized", true)
	sess.Save()
	c.JSON(message)
	return c.SendStatus(200)
}

func HandleLogout(c *fiber.Ctx) error {
	// Getting current session info
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	// Storing the current user's name
	username := sess.Get("Username")
	fmt.Println(username)

	// Deleting the cookie, (logging out)
	destroy := sess.Destroy()
	if err != nil {
		panic(destroy)
	}
	message := messages.SuccessMessage{
		Status:  "Success",
		Message: "User " + fmt.Sprint(username) + " has successfully logged out",
	}
	c.JSON(message)
	return c.SendStatus(400)
}
func HandleDeleteAccount(c *fiber.Ctx) error {
	// Getting current session info
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// checking auth
	if auth := sess.Get("Authorized"); auth == false {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	// Retrieving current user's ID
	userSessID := sess.Get("UserID")

	// id passed in param
	userID := c.Params("id")

	// if the passed id doesn't match the session stored id (user trying to delete someone else's account)
	if userSessID != userID {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User does not own that account !",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	userName := sess.Get("Username").(string)
	if deleteUser := controllers.DeleteAccount(userID); deleteUser != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Invalid user",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	message := messages.SuccessMessage{
		Status:  "Success",
		Message: "User " + userName + " has deleted their account",
	}
	c.JSON(message)
	return c.SendStatus(200)

}
func HandleGetLocation(c *fiber.Ctx) error {

	// Attempting to get current session information
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	//Checking if the current user is authorized (logged in)
	if sess.Get("Authorized") == false {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// Retrieving the ip address, and passing it to GetLocation func
	ip := c.Get("X-Forwarded-For")
	fmt.Println(ip)
	response, err := controllers.GetLocation(ip)
	if err != nil || response.Status != "success" {
		message := messages.ErrorMessage{
			Status:  "Error",
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
func HandleGetAllUsers(c *fiber.Ctx) error {
	// Getting current session info
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	fmt.Println(sess.Keys())
	// checking auth
	if auth := sess.Get("Authorized"); auth == false {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	getAllUsers, err := controllers.GetAllUsers()
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Method encountered an error",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	success := messages.AllUsersMessage{
		Status:  "Success",
		Message: "Successfully retrieved users",
		Users:   getAllUsers,
	}
	c.JSON(success)
	return c.SendStatus(200)
}
func HandleGetCurrentUser(c *fiber.Ctx) error {
	// Getting current session info
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	// checking auth
	if auth := sess.Get("Authorized"); auth == false {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	fmt.Println(sess.Get("UserID"))
	if sess.Get("UserID") == nil {
		c.WriteString("userID nill")
		return c.SendStatus(400)
	}
	currentUserID := sess.Get("UserID")

	// getting info on the current user
	getCurrentUser, err := controllers.GetCurrentUser(fmt.Sprint(currentUserID))
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Method encountered an error",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	success := messages.SuccessMessage{
		Status:  "Success",
		Message: "Successfully retrieved user",
		User:    getCurrentUser,
	}
	c.JSON(success)
	return c.SendStatus(200)
}
