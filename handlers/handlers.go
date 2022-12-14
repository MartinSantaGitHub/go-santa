package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"routes/greetings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
}

/* Handler that set the PORT and run the service */
func Handlers() {
	router := mux.NewRouter()

	// Register Home page service
	router.HandleFunc("/", home)

	// Register Greetings endpoints
	greetings.Greet(router)
	greetings.GetNames(router)

	PORT := os.Getenv("PORT")
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), handler))
}
