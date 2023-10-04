package main

import (
	"foodSharer/database"
	"foodSharer/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	app := fiber.New()
	app.Use(limiter.New())
	app.Use(logger.New())

	app.Post("/foodSharing/v1/users/signup", handlers.HandleSignUp)
	app.Post("/foodSharing/v1/users/login", handlers.HandleLogin)
	app.Post("/foodSharing/v1/users/logout", handlers.HandleLogout)
	app.Get("/foodSharing/v1/users/location", handlers.HandleGetLocation)
	app.Delete("/foodSharing/v1/users/delete/:id", handlers.HandleDeleteAccount)

	//port := os.Getenv("PORT")
	listen := app.Listen(":8080")
	if listen != nil {
		log.Println(listen)
	}

}
