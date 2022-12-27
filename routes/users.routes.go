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

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	params := mux.Vars(r)
	fmt.Println(params["id"])

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("User not found")
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := service.CreateUser(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	params := mux.Vars(r)
	fmt.Println(params["id"])

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("User not found")
		return
	}

	//Physycal delete
	db.DB.Unscoped().Delete(&user)

	//Logical Delete
	//db.DB.Delete(&user)

	w.WriteHeader(http.StatusOK)
}
