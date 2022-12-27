package service

import (
	"errors"
	"log"

	"github.com/carloscontrerasruiz/rest-api/db"
	"github.com/carloscontrerasruiz/rest-api/models"
)

func CreateUser(user *models.User) error {

	//validations
	if len(user.Firstname) <= 0 || len(user.Lastname) <= 0 {
		return errors.New("the firstname or lastname are invalids")
	}

	createdUser := db.DB.Create(&user)

	if createdUser.Error != nil {
		log.Println(createdUser.Error.Error())
		return createdUser.Error
	}

	return nil
}
