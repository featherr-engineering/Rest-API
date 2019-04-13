package utils

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 1XX Main Errors

var BadRequest = &Error{
	Code:    "400",
	Message: "Missing Headers",
}

var MissingParameters = &Error{
	Code:    "101",
	Message: "Missing Parameters",
}

var OffsetLimit = &Error{
	Code:    "102",
	Message: "Invalid offset or limit",
}

var RateLimit = &Error{
	Code:    "103",
	Message: "You exceeded the limit of requests per minute, Please try again after sometime.",
}

// 2XX

var Unauthorized = &Error{
	Code:    "200",
	Message: "Missing Headers",
}

var ErrorCodes = map[string]string{
	"100": "App Server Error",
	"101": "Missing Headers",
	"102": "Missing Parameters",
	"103": "You exceeded the limit of requests per minute, Please try again after sometime.",

	//HTTP Errors
	"110": "Un",
}

type ErrorResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Message(status int, message string) map[string]interface{} {

	resp := map[string]interface{}{
		"data":  nil,
		"error": nil,
	}

	if status > 204 {
		resp["error"] = map[string]interface{}{
			"status":  status,
			"message": message,
		}
	} else {
		resp["message"] = message
	}

	return resp
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	if data["error"] != nil {
		respError := data["error"].(map[string]interface{})
		status, _ := (respError)["status"].(int)

		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
