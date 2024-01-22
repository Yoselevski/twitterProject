package user

import (
	"twitter_app/twitter"
)

type User struct {
	UserName          string
	Followers         []string
	NumberOfFollowers int
	Following         []string
	UserTweets        []twitter.Tweet
	Feed              []twitter.Tweet
}

func CreateUser(UserName string) (*User, error) {
	user := &User{
		UserName:          UserName,
		Followers:         []string{},
		NumberOfFollowers: 0,
		Following:         []string{},
		UserTweets:        []twitter.Tweet{},
		Feed:              []twitter.Tweet{},
	}
	return user, nil
}
