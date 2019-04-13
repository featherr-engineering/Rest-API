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

//GetComments of post
var GetComments = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	data, err := models.GetComments(string(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.Respond(w, u.Message(http.StatusNotFound, "Comments not found"))
			return
		} else {
			u.Respond(w, u.Message(http.StatusInternalServerError, "Internal server error"))
			return
		}
	}

	resp := u.Message(http.StatusCreated, "Found comments")

	resp["data"] = data
	u.Respond(w, resp)
}

//CreateComment (new comment record)
var CreateComment = func(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(*models.Token)
	id := token.UserId
	comment := &models.Comment{}
	comment.UserID = id

	err := json.NewDecoder(r.Body).Decode(comment) //decode the request body into struct and failed if any error occur
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.WithFields(log.Fields{"Err": err}).Error("Could not parse request body as json")
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	if validResp, valid := comment.Validate(); !valid {
		u.Respond(w, validResp)
		return
	}

	resp := comment.Create()

	u.Respond(w, resp)
}
