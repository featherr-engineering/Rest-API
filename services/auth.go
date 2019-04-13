package services

import (
	"context"
	"github.com/featherr-engineering/rest-api/config"
	"github.com/featherr-engineering/rest-api/models"
	u "github.com/featherr-engineering/rest-api/utils"
	"github.com/getsentry/raven-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ErrorMessage(w http.ResponseWriter, message string) {
	response := make(map[string]interface{})
	response = u.Message(http.StatusForbidden, message)
	w.Header().Add("Content-Type", "application/json")
	u.Respond(w, response)
}

var TestContext = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info("Last Middleware")
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	raven.ClearContext()

	cfg := config.GetConfig()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/users/new", "/users/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                         //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			ErrorMessage(w, "Missing auth token")
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			ErrorMessage(w, "Invalid/Malformed auth token")
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			ErrorMessage(w, "Malformed authentication token")
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			ErrorMessage(w, "Token is not valid.")
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "token", tk)
		r = r.WithContext(ctx)
		h := raven.NewHttp(r)
		raven.SetHttpContext(h)

		u := &raven.User{
			ID: tk.UserId,
		}

		raven.SetUserContext(u)

		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
