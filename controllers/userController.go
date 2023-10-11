package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"foodSharer/database"
	"foodSharer/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"strings"
)

// Response Format for Geolocation api call
type Response struct {
	Status      string `json:"status"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	State       string `json:"regionName"`
	Zip         string `json:"zip"`
	Region      string `json:"region"`
}

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
func GetLocation(ip string) (Response, error) {

	url := "http://ip-api.com/json/" + ip
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}
func DeleteAccount(userID string) error {
	var userToDelete models.User

	// Create a map with "id" as the key and userID as the value
	where := map[string]interface{}{"id": userID}

	// Use the map in the Where clause to find the user
	if err := database.DB.Where(where).First(&userToDelete).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with ID %s not found", userID)
		}
		return err
	}

	if delAcc := database.DB.Delete(&userToDelete).Error; delAcc != nil {
		return delAcc
	}
	return nil
}
func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if findAllUsers := database.DB.Find(&users).Error; findAllUsers != nil {
		return []*models.User{}, findAllUsers
	}
	return users, nil
}
func GetCurrentUser(userID string) (*models.User, error) {
	var user *models.User
	if findAllUsers := database.DB.First(&user,userID).Error; findAllUsers != nil {
		return nil, findAllUsers
	}
	return user, nil
}
