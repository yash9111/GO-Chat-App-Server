package main

import (
	"fmt"
	"net/http"
	"os"
	"server/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sendOtp/", routes.SentOtpEndpoint).Methods("POST")
	r.HandleFunc("/verify/", routes.VerifyOtpEndpoint).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port if not set
	}
	fmt.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
