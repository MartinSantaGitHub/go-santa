package middlewares

import (
	"net/http"
)

func ValidateContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnprocessableEntity)

			return
		}

		next.ServeHTTP(w, r)
	}
}
