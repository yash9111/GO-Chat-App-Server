package routes

import (
	"encoding/json"
	"net/http"
	"server/helpers"
	"server/vars"
)

func GetAllChatsEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var userDetails vars.UserDetails
	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	allChats := helpers.GetAllChats(userDetails.MobileNumber)

	json.NewEncoder(w).Encode(allChats)
}
