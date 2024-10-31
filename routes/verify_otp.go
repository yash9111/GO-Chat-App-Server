package routes

import (
	"encoding/json"
	"net/http"
	"server/vars"
)

func VerifyOtpEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var varOtp vars.VerifyOtp

	err := json.NewDecoder(r.Body).Decode(&varOtp)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if vars.Otps[varOtp.MobileNumber] == varOtp.Otp {
		json.NewEncoder(w).Encode(vars.Verify{MobileNumber: varOtp.MobileNumber, Msg: "OK"})
	} else {
		json.NewEncoder(w).Encode(vars.Verify{MobileNumber: varOtp.MobileNumber, Msg: "NOT_OK"})
	}

}
