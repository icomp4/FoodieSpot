package main

import (
	"foodSharer/database"
	"foodSharer/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	app := fiber.New()
	app.Use(limiter.New())
	app.Use(logger.New())

	app.Post("/foodSharing/v1/users/signup", handlers.HandleSignUp)              // user signup
	app.Post("/foodSharing/v1/users/login", handlers.HandleLogin)                // user login
	app.Post("/foodSharing/v1/users/logout", handlers.HandleLogout)              // user logout
	app.Get("/foodSharing/v1/users/location", handlers.HandleGetLocation)        // retrieve the current user's location
	app.Get("/foodSharing/v1/user", handlers.HandleGetCurrentUser)               // return a list of all users
	app.Get("/foodSharing/v1/users/all", handlers.HandleGetAllUsers)             // get the current user's details
	app.Delete("/foodSharing/v1/users/delete/:id", handlers.HandleDeleteAccount) // delete current user
	app.Post("foodSharing/v1/post/create",handlers.HandleCreatePost) // Create a new post
	app.Get("foodSharing/v1/post/:id",handlers.HandleFetchPost)

	//port := os.Getenv("PORT")
	listen := app.Listen(":8080")
	if listen != nil {
		log.Println(listen)
	}

}
