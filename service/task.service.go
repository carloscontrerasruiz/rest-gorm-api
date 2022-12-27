package service

import (
	"errors"
	"fmt"

	"github.com/carloscontrerasruiz/rest-api/db"
	"github.com/carloscontrerasruiz/rest-api/models"
)

func CreateTask(task *models.Task) error {

	//validations
	if len(task.Title) <= 0 || len(task.Descripton) <= 0 {
		return errors.New("the title or description are invalids")
	}

	createdTask := db.DB.Create(&task)

	if createdTask.Error != nil {
		fmt.Println(createdTask.Error.Error())
		return createdTask.Error
	}

	return nil
}
