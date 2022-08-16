package main

import (
	"github/ctatis1/gorm-restapi/database"
	"github/ctatis1/gorm-restapi/models"
	"github/ctatis1/gorm-restapi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.DBConnection()

	database.DB.AutoMigrate(models.User{})
	database.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}