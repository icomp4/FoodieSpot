package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username          string
	Password          string
	Followers         []*User     `gorm:"many2many:user_followers;"`
	Following         []*User     `gorm:"many2many:user_following;"`
	FavoriteLocations []*Location `gorm:"many2many:user_favorite_locations;"`
	FollowerCount     int
	FollowingCount    int
}

type Location struct {
	gorm.Model
	Name        string
	Address     string
	Rating      float32
	ImageURL    string
	Description string
	Category    string
	Latitude    float64
	Longitude   float64
}
