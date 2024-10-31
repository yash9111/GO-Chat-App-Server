package vars

var Otps = make(map[string]string)

type UserDetails struct {
	MobileNumber string `json:"mobile"`
}
type VerifyOtp struct {
	MobileNumber string `json:"mobile"`
	Otp          string `json:"otp"`
}

type Verify struct {
	MobileNumber string `json:"mobile"`
	Msg          string `json:"msg"`
}

type TwilloResponse struct {
	Status string `json:"status"`
}
