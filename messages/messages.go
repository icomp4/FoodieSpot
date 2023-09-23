package messages

import "foodSharer/models"

type SuccessMessage struct {
	Status  string
	Message string
	User    *models.User
}
type ErrorMessage struct {
	Status  string
	Message string
}
