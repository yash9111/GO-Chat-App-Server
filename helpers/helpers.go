package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"server/vars"
	"strings"

	"github.com/joho/godotenv"
)

func GenerateOtp() string {
	return "289328"
}

func SendMessage(mobile string, Otps map[string]string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading environment variables: " + err.Error())
	}

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	log.Println("Account SID:", accountSid)
	log.Println("Auth Token:", authToken)

	apiUrl := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	otp := GenerateOtp()
	Otps[mobile] = otp

	formData := url.Values{}
	formData.Set("To", "+91"+mobile)
	formData.Set("From", "+18594792474") // Replace with your Twilio number
	formData.Set("Body", "Hello, here is your OTP for WhatsApp login: "+otp)

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

	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Println("Twilio Response Body:", string(bodyBytes))
	resp.Body = io.NopCloser(bytes.NewReader(bodyBytes)) // Reassigning to resp.Body for decoding

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error from Twilio: Status Code: %d\nResponse: %s\n", resp.StatusCode, bodyBytes)
		return "error" // Handle error accordingly
	}

	var twilioResponse vars.TwilloResponse
	err = json.NewDecoder(resp.Body).Decode(&twilioResponse)
	if err != nil {
		log.Printf("Error decoding Twilio response: %v\n", err)
		return "error" // Handle error accordingly
	}

	return twilioResponse.Status
}
