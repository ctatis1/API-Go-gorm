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

	http.ListenAndServe(":3000", router)
}