package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carloscontrerasruiz/rest-api/db"
	"github.com/carloscontrerasruiz/rest-api/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUsersHandler"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUserHandler"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(r.Body).Decode(&user)

	//validations
	if len(user.Firstname) <= 0 || len(user.Lastname) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUser := db.DB.Create(&user)

	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(createdUser.Error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUserHandler"))
}
