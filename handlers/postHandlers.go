package handlers

import (
	"encoding/json"
	"fmt"
	"foodSharer/controllers"
	"foodSharer/messages"
	"foodSharer/models"

	"github.com/gofiber/fiber/v2"
)

// insane spaghetti code (yikes)
func HandleCreatePost(c *fiber.Ctx) error {
	var post models.Post
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	if auth := sess.Get("Authorized"); auth != true {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User is not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	currUserID := sess.Get("UserID")
	curUser, err := controllers.GetCurrentUser(fmt.Sprint(currUserID))
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Unable to find specified user",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}

	if enc := json.Unmarshal(c.Body(), &post); enc != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Failed to unmarshal response body",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	// Convert the string to a uint
	post.AuthorID = curUser.ID
	post.Author = *curUser
	if createPost := controllers.CreatePost(fmt.Sprint(currUserID), post); createPost != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Failed to create post",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	message := messages.SuccessfulPostCreation{
		Status:  "Success",
		Message: "Post has successfully been created",
		Likes:   post.Likes,
	}
	
	message.Post.Author.UserID = fmt.Sprint(curUser.ID)
	message.Post.Author.Username = curUser.Username
	
	message.Post.Location.Name = post.Location.Name
	message.Post.Location.Address = post.Location.Address
	message.Post.Location.Rating = post.Location.Rating
	message.Post.Location.ImageURL = post.Location.ImageURL
	message.Post.Location.Description = post.Location.Description
	message.Post.Location.Category = post.Location.Category
	message.Post.Location.Latitude = post.Location.Latitude
	message.Post.Location.Longitude = post.Location.Longitude
	c.JSON(message)
	return c.SendStatus(201)
}
func HandleFetchPost(c *fiber.Ctx) error {
	postID := c.Query("id")
	sess, err := Store.Get(c)
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Session Invalid",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	if auth := sess.Get("Authorized"); auth != true {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "User is not authorized",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	post, err := controllers.RetrievePost(postID) 
	if err != nil {
		message := messages.ErrorMessage{
			Status:  "Error",
			Message: "Failed to fetch post with that ID",
		}
		c.JSON(message)
		return c.SendStatus(400)
	}
	message := messages.SuccessPostFetch{
		Status:  "Success",
		Message: "Successfully fetched post",
		Likes: post.Likes,
	}
	message.Post.Author.UserID = fmt.Sprint(post.Author.ID)
	message.Post.Author.Username = post.Author.Username
	
	message.Post.Location.Name = post.Location.Name
	message.Post.Location.Address = post.Location.Address
	message.Post.Location.Rating = post.Location.Rating
	message.Post.Location.ImageURL = post.Location.ImageURL
	message.Post.Location.Description = post.Location.Description
	message.Post.Location.Category = post.Location.Category
	message.Post.Location.Latitude = post.Location.Latitude
	message.Post.Location.Longitude = post.Location.Longitude
	c.JSON(message)
	return c.SendStatus(200)

}