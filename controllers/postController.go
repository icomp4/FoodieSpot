package controllers

import (
	"foodSharer/database"
	"foodSharer/models"
)

func CreatePost(userID string, post models.Post) (models.Post, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return models.Post{}, err
	}
	if create := database.DB.Create(post).Error; create != nil {
		return models.Post{}, create
	}
	return post, nil
}
