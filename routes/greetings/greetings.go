package greetings

import (
	"controllers/greetings"
	"helpers"
	"middlewares"

	"github.com/gorilla/mux"
)

/* Greet greets an user */
func Greet(router *mux.Router) {
	router.HandleFunc("/hello", helpers.MultipleMiddleware(greetings.Greet, middlewares.CheckDB, middlewares.ValidateName)).Methods("POST")
}
