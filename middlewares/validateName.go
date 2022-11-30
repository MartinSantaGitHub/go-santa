package middlewares

import (
	"context"
	"encoding/json"
	"helpers"
	mr "models/request"
	"net/http"
)

func ValidateName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user mr.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "There was an error receiving the data: "+err.Error(), 400)

			return
		}

		if len(user.Name) == 0 {
			http.Error(w, "The name is required", 400)

			return
		}

		ctx := context.WithValue(r.Context(), helpers.RequestUserKey{}, user)
		r = r.Clone(ctx)

		next.ServeHTTP(w, r)
	}
}
