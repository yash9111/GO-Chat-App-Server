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
	r.HandleFunc("/getAllChats/", routes.GetAllChatsEndpoint).Methods("POST")
	r.HandleFunc("/ws", routes.ServeWs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
