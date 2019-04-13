package main

import (
	"fmt"
	"github.com/featherr-engineering/rest-api/config"
	"github.com/featherr-engineering/rest-api/controllers"
	"github.com/featherr-engineering/rest-api/services"
	"github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"os"

	"net/http"
)

var cfg = config.GetConfig()

func init() {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	raven.SetEnvironment(cfg.AppEnv)
	raven.SetRelease("0.1.0")

	err := raven.SetDSN("https://d213ea30b5b54158a27dece2072d1c8f:788a772b517642e08cbd683c3d3b1602@sentry.io/1438262")

	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	router := mux.NewRouter()

	router.Use(services.JwtAuthentication)
	router.Use(raven.Recoverer)

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
		raven.CaptureErrorAndWait(err, nil)
		log.Panic(err)
	}
}
