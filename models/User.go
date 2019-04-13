package models

import (
	"github.com/featherr-engineering/rest-api/config"
	u "github.com/featherr-engineering/rest-api/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//a struct to rep user
type User struct {
	GormModel
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Subscriptions string `json:"subscriptions"`
	Image         string `json:"image"`
	FcmToken      string `json:"fcmToken"`
}

/*
JWT claims struct
*/
type Token struct {
	UserId string
	jwt.StandardClaims
}

var cfg = config.GetConfig()

func (user *User) Validate() *u.APIError {
	Api := &u.APIError{
		Code: http.StatusUnprocessableEntity,
	}

	if !strings.Contains(user.Email, "@") {
		Api.Message = "Email address is required"
		return Api
	}

	if len(user.Password) < 6 {
		Api.Message = "Password is required"
		return Api
	}

	temp := &User{}

	err := GetDB().Table("users").Where("email = ?", user.Email).Or("username = ?", user.Username).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		Api.Message = "Connection error. Please retry"
		return Api
	}

	if temp.Username != "" && temp.Username == user.Username {
		Api.Message = "Username is already in use by another user."
		return Api
	}

	if temp.Email != "" && temp.Email == user.Email {
		Api.Message = "Email address already in use by another user."
		return Api
	}

	return nil
}

func (user *User) Create() map[string]interface{} {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err := GetDB().Create(user).Error
	if err != nil {
		return u.Message(http.StatusBadRequest, "Could not create user")
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(cfg.JWTSecret))
	//user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message(http.StatusOK, "User has been created")
	response["data"] = map[string]interface{}{
		"user":  user,
		"token": tokenString,
	}

	return response
}

func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(http.StatusUnprocessableEntity, "Email address not found")
		}
		return u.Message(http.StatusUnprocessableEntity, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(http.StatusUnprocessableEntity, "Invalid login credentials. Please try again")
	}

	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(cfg.JWTSecret))
	//user.Token = tokenString //Store the token in the response

	resp := u.Message(http.StatusOK, "Logged In")
	resp["data"] = map[string]interface{}{
		"user":  user,
		"token": tokenString,
	}

	return resp
}

func GetUser(u uint) *User {

	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Email == "" { //User not found!
		return nil
	}

	user.Password = ""
	return user
}
