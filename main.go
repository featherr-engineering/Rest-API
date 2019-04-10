package main

import (
	"fmt"
	"github.com/abdullahi/feather-backend/controllers"
	"github.com/abdullahi/feather-backend/services"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(services.JwtAuthentication)

	router.HandleFunc("/users/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/login", controllers.Authenticate).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
