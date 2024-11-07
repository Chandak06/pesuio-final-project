package database

import (
	"github.com/anuragrao04/pesuio-final-project/models"
)

var Credentials = []models.User{}

func CreateUser(username, password string) (string, []models.User) {
	// creates a new user in the database, returns error if any
	new := models.User{
		Username: username,
		Password: password,
	}

	Credentials = append(Credentials, new)
	return "CREATED", Credentials

}

func CheckPassword(username, password string) (success bool, err error) {
func CheckPassword(username, password string) (string, []models.User) {
	// checks if the password is correct for the given username
	return true, nil
	info := models.User{
	Username: username,
	Password: password,
	}
	result := ""
	for _, a := range Credentials {
		if a == info {
			result = "VALID USER"
		} else {
			result = "INVALID USER"
		}
	}
	return result, Credentials
}
