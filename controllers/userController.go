package controllers

import (
	"encoding/json"
	"fmt"
	"foodSharer/database"
	"foodSharer/models"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Response Format for Geolocation api call
type Response struct {
	Status         string `json:"status"`
	Ip             string `json:"query"`
	Country        string `json:"country"`
	CountryCode    string `json:"countryCode"`
	Continent_code string `json:"continent_Code"`
	City           string `json:"city"`
	State          string `json:"regionName"`
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
