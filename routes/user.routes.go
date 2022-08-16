package routes

import (
	"encoding/json"
	"github/ctatis1/gorm-restapi/database"
	"github/ctatis1/gorm-restapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return 
	}

	database.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := database.DB.Create(&user)	
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))	
	}

	json.NewEncoder(w).Encode(&user)
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("User not found"))
		return
	}

	database.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)

}