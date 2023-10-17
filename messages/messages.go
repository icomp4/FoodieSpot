package messages

import (
	"foodSharer/controllers"
	"foodSharer/models"
)

type SuccessMessage struct {
	Status  string
	Message string
	User    *models.User
}
type ErrorMessage struct {
	Status  string
	Message string
}

type LocationMessage struct {
	Status   string
	Message  string
	Location controllers.Response
}
type AllUsersMessage struct {
	Status  string
	Message string
	Users   []*models.User
}
type SuccessfulPostCreation struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Post    struct {
		Author struct {
			UserID   string
			Username string
		}
		Location struct {
			Name        string
			Address     string
			Rating      float32
			ImageURL    string
			Description string
			Category    string
			Latitude    float64
			Longitude   float64
		}
	}
	Likes int `json:"likes"`
}

type SuccessPostFetch struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Post    struct {
		Author struct {
			UserID   string
			Username string
		}
		Location struct {
			Name        string
			Address     string
			Rating      float32
			ImageURL    string
			Description string
			Category    string
			Latitude    float64
			Longitude   float64
		}
	}
	Likes int `json:"likes"`
}
