package helpers

import (
	"server/vars"
	"time"
)

func GetAllChats(user string) []vars.ChatPreview {
	return []vars.ChatPreview{
		{User: vars.User{
			Name:         "Yash",
			MobileNumber: "382332332",
			ImageUrl:     "",
		},
			ChatId: "32udia3",
			LastMsg: vars.Message{
				Time:    time.Now(),
				Content: "Hi",
				SentBy:  "9111607983",
				Status:  "SENT",
			},
		},
		{
			User: vars.User{
				Name:         "Pooja",
				MobileNumber: "91282923",
				ImageUrl:     "",
			},
			ChatId: "sua291jiawu",
			LastMsg: vars.Message{
				Time:    time.Now(),
				Content: "Hello",
				SentBy:  "7771977187",
				Status:  "VIEWED",
			},
		},
	}
}
