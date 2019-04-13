package models

import (
	u "github.com/featherr-engineering/rest-api/utils"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	TBH                string = "#TBH"
	CRUSHES            string = "UNO CRASHES"
	GAMING             string = "Gaming"
	Fitness            string = "Fitness"
	Sports             string = "Sports"
	Music              string = "Music"
	Movies             string = "Movies"
	News               string = "News"
	Conservative       string = "Conservative"
	Liberal            string = "Liberal"
	Business           string = "Business"
	ScienceEngineering string = "Science and Engineering"
	Stories            string = "Stories"
	Anonymous          string = "Anonymous"
)

type Post struct {
	GormModel
	Text          string `json:"text"`
	Category      string `gorm:"type:enum('#TBH','UNO CRASHES','Gaming','Fitness', 'Sports', 'Music','Movies','News','Conservative','Liberal','Business','Science and Engineering','Stories','Anonymous'); not null" json:"category"`
	VotesCount    int    `json:"votesCount"`
	CommentsCount int    `json:"commentsCount"`
	Time          string `json:"time" sql:"DEFAULT:'0'"`
	Color         string `json:"color"`
	Image         string `json:"image" sql:"DEFAULT:'null'"`
	User          User   `json:"user"`
	UserID        string `json:"-"`
	Votes         []Vote `json:"votes"`
}

func (post *Post) Validate() *u.APIError {
	Api := &u.APIError{
		Code: http.StatusUnprocessableEntity,
	}

	if len(post.Text) <= 5 {
		Api.Message = "text field must be at least 6 characters long"
		return Api
	} else if len(post.Category) <= 0 {
		Api.Message = "category field must be at least 1 characters long"
		return Api
	} else {
		return nil
	}
}

func (post *Post) Create() *u.APIError {

	validErr := post.Validate()

	if validErr != nil {
		return validErr
	}

	post.Color = "#ffffff"

	err := GetDB().Create(post).Related(&post.User).Error

	if err != nil {
		return &u.APIError{
			Code:    http.StatusBadRequest,
			Message: "Could not add comment",
		}
	} else {
		return nil
	}
}

func GetPost(id string) (*Post, error) {
	post := &Post{}
	err := GetDB().Table("posts").Preload("User").Where("id = ?", id).First(post).Error

	return post, err
}

func GetPosts() []*Post {

	posts := make([]*Post, 0)

	err := GetDB().Table("posts").Preload("User").Order("LOG10(ABS(votes_count) + 1) * SIGN(votes_count) + (UNIX_TIMESTAMP(created_at)/300000) desc", true).Limit(100).Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		raven.CaptureErrorAndWait(err, nil)
		log.WithFields(log.Fields{"Err": err}).Error("Could not create post")
		return nil
	}

	return posts
}
