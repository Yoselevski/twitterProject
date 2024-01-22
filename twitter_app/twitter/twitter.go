package twitter

import (
	"time"
)


type Tweet struct {
	UserName string
	Content string
	Date    time.Time
}