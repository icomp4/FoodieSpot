package main

import (
	"foodSharer/database"
	"foodSharer/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	err := database.StartDB()
	if err != nil {
		log.Println(err)
	}
	app := fiber.New()

	app.Post("/foodSharing/v1/users/signup", handlers.HandleSignUp)
	app.Post("/foodSharing/v1/users/login", handlers.HandleLogin)

	listen := app.Listen(":8080")
	if listen != nil {
		log.Println(listen)
	}

}
