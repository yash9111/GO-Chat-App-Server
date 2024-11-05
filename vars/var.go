package vars

import (
	"time"
)

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

type User struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile"`
	ImageUrl     string `json:"img_url"`
}

type Message struct {
	Time    time.Time `json:"time"`
	Content string    `json:"content"`
	SentBy  string    `json:"sent_by"`
	Status  string    `json:"status"`
	Id      string    `json:"msg_id"`
}

type ChatPreview struct {
	User    User    `json:"user"`
	ChatId  string  `json:"id"`
	LastMsg Message `json:"last_msg"`
}

type UserChats struct {
	ChatId string    `json:"id"`
	Chats  []Message `json:"chats"`
}
