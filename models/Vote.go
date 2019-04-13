package models

import (
	"fmt"
	u "github.com/featherr-engineering/rest-api/utils"
	"net/http"
)

type Vote struct {
	GormModel
	Dir    int    `json:"dir"`
	Post   Post   `json:"-"`
	PostId string `json:"postId"`
	User   User   `json:"-"`
	UserID string `json:"userId"`
}

//func (post *Vote) Validate() (map[string]interface{}, bool) {
//	if len(post.Text) <= 0 {
//		return u.Message(http.StatusUnprocessableEntity, "Post should contain text"), false
//	}
//
//	return u.Message(http.StatusOK, "success"), true
//}

func (vote *Vote) Create() map[string]interface{} {
	//if resp, ok := post.Validate(); !ok {
	//	return resp
	//}

	GetDB().Create(vote).Related(&vote.Post).Related(&vote.User)

	post, _ := GetPost(string(vote.PostId))

	GetDB().Model(&post).Update("votes_count", post.VotesCount+vote.Dir)

	response := u.Message(http.StatusCreated, "Vote has been submitted")

	response["data"] = vote

	return response
}

func GetVote(id string) (*Vote, error) {
	vote := &Vote{}
	err := GetDB().Table("votes").Preload("User").Where("id = ?", id).First(vote).Error

	return vote, err
}

func GetVotes() []*Vote {
	votes := make([]*Vote, 0)

	err := GetDB().Table("votes").Preload("User").Find(&votes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return votes
}
