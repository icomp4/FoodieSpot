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
func Login(username string, password string) (*models.User, error) {
	var user *models.User

	// Indexing the database for a user with the given username
	err := database.DB.Where("Username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	//comparing the given password to the hash stored for the specified user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
