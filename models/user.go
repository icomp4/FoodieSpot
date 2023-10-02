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
	Posts             []Post `gorm:"foreignKey:AuthorID"` // This field represents the posts authored by the user
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

type Post struct {
	gorm.Model
	AuthorID   uint      // This field will store the ID of the author user
	Author     User      `gorm:"foreignKey:AuthorID"` // This field represents the author user
	LocationID uint      // This field will store the ID of the location
	Location   *Location `gorm:"foreignKey:LocationID"` // This field represents the location
	Likes      int
}
