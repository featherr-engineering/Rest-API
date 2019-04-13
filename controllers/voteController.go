package controllers

import (
	"encoding/json"
	"github.com/featherr-engineering/rest-api/models"
	u "github.com/featherr-engineering/rest-api/utils"
	"github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//Get one vote record
var GetVote = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	data, err := models.GetVote(string(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.Respond(w, u.Message(http.StatusNotFound, "Vote not found"))
			return
		} else {
			u.Respond(w, u.Message(http.StatusInternalServerError, "Internal server error"))
			return
		}
	}

	resp := u.Message(http.StatusCreated, "Found vote")

	resp["data"] = data
	u.Respond(w, resp)
}

//Get all vote records
var GetAllVotes = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetVotes()

	resp := u.Message(http.StatusCreated, "Found vote")

	resp["data"] = data
	u.Respond(w, resp)
}

//Create one vote record
var CreateVote = func(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(*models.Token)
	id := token.UserId
	vote := &models.Vote{}
	vote.UserID = id

	err := json.NewDecoder(r.Body).Decode(vote) //decode the request body into struct and failed if any error occur
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.WithFields(log.Fields{"Err": err}).Error("Could not parse request body as json")
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	resp := vote.Create()

	u.Respond(w, resp)
}
