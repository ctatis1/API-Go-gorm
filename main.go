package main

import (
	"github/ctatis1/gorm-restapi/database"
	"github/ctatis1/gorm-restapi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.DBConnection()

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}