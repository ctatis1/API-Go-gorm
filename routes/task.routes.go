package routes

import (
	"encoding/json"
	"github/ctatis1/gorm-restapi/database"
	"github/ctatis1/gorm-restapi/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	database.DB.Find(&tasks)

	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	database.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return 
	}
	json.NewEncoder(w).Encode(&task)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := database.DB.Create(&task)
	if err := createdTask.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	database.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
	}

	database.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}