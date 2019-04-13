package controllers

import (
	"encoding/json"
	"github.com/abdullahi/feather-backend/models"
	u "github.com/abdullahi/feather-backend/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

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

var CreateComment = func(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(*models.Token)
	id := token.UserId
	comment := &models.Comment{}
	comment.UserID = id

	err := json.NewDecoder(r.Body).Decode(comment) //decode the request body into struct and failed if any error occur
	if err != nil {
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
