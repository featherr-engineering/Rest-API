package models

import (
	"github.com/featherr-engineering/rest-api/utils"
	"net/http"
)

type Comment struct {
	GormModel
	Text       string `json:"text"`
	Time       string `json:"time"`
	VotesCount int    `json:"votesCount"`
	User       User   `json:"user"`
	UserID     string `json:"userId"`
	PostID     string `json:"postId"`
}

func GetComments(id string) ([]*Comment, error) {
	comments := make([]*Comment, 0)
	err := GetDB().Table("comments").Preload("User").Where("post_id = ?", id).Order("created_at DESC").Find(&comments).Error

	return comments, err
}

func (comment *Comment) Validate() (map[string]interface{}, bool) {
	if len(comment.Text) <= 5 {
		return utils.Message(http.StatusUnprocessableEntity, "text field must be at least 6 characters long"), false
	} else if len(comment.PostID) <= 0 {
		return utils.Message(http.StatusUnprocessableEntity, "invalid post id reference"), false
	} else {
		return nil, true
	}
}

func (comment *Comment) Create() map[string]interface{} {
	post, _ := GetPost(comment.PostID)

	GetDB().Model(&post).Update("comments_count", post.CommentsCount+1)
	err := GetDB().Create(comment).Error

	if err != nil {
		return utils.Message(http.StatusBadRequest, "Could not add comment")
	}

	resp := utils.Message(http.StatusOK, "Comment has been created")

	resp["data"] = post

	return resp

}
