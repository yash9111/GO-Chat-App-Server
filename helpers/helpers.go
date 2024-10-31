package helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"server/vars"
	"strings"

	"github.com/joho/godotenv"
)

func GenerateOtp() string {

	return string("289328")

}
func SendMessage(mobile string, Otps map[string]string) string {

	err := godotenv.Load()

	if err != nil {
		log.Panic("Something went srong " + string(err.Error()))
	}
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	// Create the request URL
	apiUrl := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	otp := GenerateOtp()
	Otps[mobile] = otp

	formData := url.Values{}
	formData.Set("To", "+91"+mobile)
	formData.Set("From", "+18594792474")
	formData.Set("Body", "Hello, here is your otp for whatsapp login "+otp)

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	defer resp.Body.Close()

	var twilloResponse vars.TwilloResponse
	err = json.NewDecoder(resp.Body).Decode(&twilloResponse)

	return twilloResponse.Status
}
