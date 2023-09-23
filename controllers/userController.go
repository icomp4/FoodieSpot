package controllers

import (
	"foodSharer/database"
	"foodSharer/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

func SignUp(user *models.User) error {
	user.Username = strings.ToLower(user.Username)
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = string(password)
	if err := database.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
