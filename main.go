package main

import (
	"fmt"
	"net/http"
	"os"
	"server/routes"

	"github.com/gorilla/mux"
)

func demo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, World!")

}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sendOtp/", routes.SentOtpEndpoint).Methods("POST")
	r.HandleFunc("/verify/", routes.VerifyOtpEndpoint).Methods("POST")
	r.HandleFunc("/", demo).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // default port if not set
	}
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
