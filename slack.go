package phoenix

import (
	"time"
)

type (
	Message struct {
		Token       string
		TeamId      string
		ChannelId   string
		ChannelName string
		Timestamp   time.Time
		UserId      string
		Username    string
		Text        string
		TriggerWord string
	}

	Response struct {
		Text     string `json:"text"`
		Username string `json:"username,omitempty"`
		Parse    string `json:"parse,omitempty"`
	}
)
