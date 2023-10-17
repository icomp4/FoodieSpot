package controllers

import (
	"foodSharer/database"
	"foodSharer/models"
)

func CreatePost(userID string, post models.Post)  error {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return err
	}
	if create := database.DB.Preload("Location").Create(&post).Error; create != nil {
		return create
	}
	return nil
}
