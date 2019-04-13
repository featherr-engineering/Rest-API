package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/abdullahi/feather-backend/models"
	u "github.com/abdullahi/feather-backend/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

var GetPost = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	data, err := models.GetPost(string(id))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.Respond(w, u.Message(http.StatusNotFound, "Post not found"))
			return
		} else {
			u.Respond(w, u.Message(http.StatusInternalServerError, "Internal server error"))
			return
		}
	}

	resp := u.Message(http.StatusCreated, "Found post")

	resp["data"] = data
	u.Respond(w, resp)
}

var GetAllPosts = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetPosts()

	resp := u.Message(http.StatusCreated, "Found post")

	resp["data"] = data
	u.Respond(w, resp)
}

var CreatePost = func(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(*models.Token)
	id := token.UserId
	post := &models.Post{}
	post.UserID = id

	err := json.NewDecoder(r.Body).Decode(post) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(http.StatusBadRequest, "Invalid request"))
		return
	}

	gormErr := post.Create()

	if gormErr != nil {
		fmt.Println(gormErr)
		u.Respond(w, u.Message(gormErr.Code, gormErr.Message))
		return
	}

	response := u.Message(http.StatusCreated, "Post has been created")

	response["data"] = post

	u.Respond(w, response)
}
