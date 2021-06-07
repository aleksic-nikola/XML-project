package data

import(
	"time"
)

type Notification struct {
	ID int `json:"id""`
	From string `json:"from"`
	For string `json:"for"`
	IsRead bool `json:"isread"`
	Timestamp time.Time `json:"timestamp"`
}