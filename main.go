package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/FranciscoJSB12/go-gorm-rest-api/db"
	"github.com/FranciscoJSB12/go-gorm-rest-api/models"
	"github.com/FranciscoJSB12/go-gorm-rest-api/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
