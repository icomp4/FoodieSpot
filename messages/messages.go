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
