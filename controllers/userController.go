package controllers

import (
	"foodSharer/database"
	"foodSharer/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

func SignUp(user *models.User) error {
	// Storing usernames in lowercase, not necessary but looks cleaner
	user.Username = strings.ToLower(user.Username)

	// Hashing password with the default bcrypt cost
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}

	// updating user password to hashed password
	user.Password = string(password)
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
