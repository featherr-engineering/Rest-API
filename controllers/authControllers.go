package controllers

import (
	"encoding/json"
	"github.com/featherr-engineering/rest-api/models"
	u "github.com/featherr-engineering/rest-api/utils"
	"github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Sign Up
var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.WithFields(log.Fields{"Err": err}).Error("Could not parse request body as json")
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	validErr := user.Validate()

	if validErr != nil {
		u.Respond(w, u.Message(validErr.Code, validErr.Message))
		return
	}

	resp := user.Create() //Create account
	u.Respond(w, resp)
}

//Login
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.WithFields(log.Fields{"Err": err}).Error("Could not parse request body as json")
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}
