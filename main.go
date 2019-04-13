package main

import (
	"fmt"
	"github.com/abdullahi/feather-backend/config"
	"github.com/abdullahi/feather-backend/controllers"
	"github.com/abdullahi/feather-backend/services"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	cfg := config.GetConfig()

	router := mux.NewRouter()

	router.Use(services.JwtAuthentication)

	router.HandleFunc("/users/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts/new", controllers.CreatePost).Methods("POST")

	router.HandleFunc("/posts/{id}/comments", controllers.GetComments).Methods("GET")
	router.HandleFunc("/comments", controllers.CreateComment).Methods("POST")

	router.HandleFunc("/vote", controllers.CreateVote).Methods("POST")

	router.HandleFunc("/vote", controllers.GetAllVotes).Methods("GET")
	router.HandleFunc("/vote/{id}", controllers.GetVote).Methods("GET")
	router.HandleFunc("/vote", controllers.CreateVote).Methods("POST")

	port := cfg.AppPort
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
