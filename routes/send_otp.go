package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/helpers"
	"server/vars"
)

func SentOtpEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var userDetails vars.UserDetails
	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	status := helpers.SendMessage(userDetails.MobileNumber, vars.Otps)

	fmt.Print("msg" + status)

	var message string

	if status == "queued" {
		message = "OK"
	} else {
		message = "NOT_OK"
	}
	json.NewEncoder(w).Encode(vars.Verify{MobileNumber: userDetails.MobileNumber,
		Msg: message,
	})
}
