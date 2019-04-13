package controllers

import (
	"encoding/json"
	"github.com/abdullahi/feather-backend/models"
	u "github.com/abdullahi/feather-backend/utils"
	"net/http"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
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

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}
