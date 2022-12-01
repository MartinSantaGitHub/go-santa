package greetings

import (
	"controllers/greetings"
	"helpers"
	"middlewares"

	"github.com/gorilla/mux"
)

/* Greet greets an user */
func Greet(router *mux.Router) {
	router.HandleFunc("/hello", helpers.MultipleMiddleware(greetings.Greet,
		middlewares.CheckDB,
		middlewares.ValidateContentType,
		middlewares.ValidateName)).Methods("POST")
}

/* GetNames gets the names */
func GetNames(router *mux.Router) {
	router.HandleFunc("/hello", helpers.MultipleMiddleware(greetings.GetNames,
		middlewares.CheckDB)).Methods("GET")
}
