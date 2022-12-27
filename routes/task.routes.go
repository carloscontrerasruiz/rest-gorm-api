package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/carloscontrerasruiz/rest-api/db"
	"github.com/carloscontrerasruiz/rest-api/models"
	"github.com/carloscontrerasruiz/rest-api/service"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println(task)

	err := service.CreateTask(&task)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task

	params := mux.Vars(r)
	fmt.Println(params["id"])

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Task not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task

	params := mux.Vars(r)
	fmt.Println(params["id"])

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Task not found")
		return
	}

	//Physycal delete
	db.DB.Unscoped().Delete(&task)

	//Logical Delete
	//db.DB.Delete(&task)

	w.WriteHeader(http.StatusNoContent)
}
