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
	Success        bool   `json:"success"`
	Ip             string `json:"ip"`
	Country        string `json:"country"`
	Country_code   string `json:"country_Code"`
	Continent      string `json:"continent"`
	Continent_code string `json:"continent_Code"`
	City           string `json:"city"`
	State          string `json:"region"`
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

	url := "https://ip-geolocation-ipwhois-io.p.rapidapi.com/json/?ip=" + ip
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "c774f6f2b9msh1127b6e1b245602p1e2e23jsn6b577b35a63f")
	req.Header.Add("X-RapidAPI-Host", "ip-geolocation-ipwhois-io.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}
	fmt.Println(response.Success)
	fmt.Println(response.City)
	return response, nil
}
