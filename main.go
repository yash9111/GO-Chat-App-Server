package main

import (
	"fmt"
	"log"
	"net/http"
	"server/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sendOtp/", routes.SentOtpEndpoint).Methods("POST")
	r.HandleFunc("/verify/", routes.VerifyOtpEndpoint).Methods("POST")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
