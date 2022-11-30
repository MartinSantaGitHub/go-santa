package greetings

import (
	"encoding/json"
	"fmt"
	"net/http"

	"db"
	"helpers"
	req "models/request"
	res "models/response"
)

// region "Actions"

/* Greet greets an user */
func Greet(w http.ResponseWriter, r *http.Request) {
	var response res.Greeting

	user := r.Context().Value(helpers.RequestUserKey{}).(req.User)

	isFound, err := db.DbConn.SaveName(user.Name)

	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if isFound {
		response = res.Greeting{
			Message: fmt.Sprintf("Hello, %s! Welcome back!", user.Name),
			Exists:  true,
		}
	} else {
		response = res.Greeting{
			Message: fmt.Sprintf("Hello, %s!", user.Name),
		}

		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(response)
}

// endregion
