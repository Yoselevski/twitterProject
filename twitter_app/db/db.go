package db

import (
	"fmt"
	"time"
	"twitter_app/helper"
	"twitter_app/twitter"
	"twitter_app/user"
)

var DatabaseUser = make(map[string]user.User)
var counterUsers = 0

// InsertUser inserts a new user into the database
func InsertUser(userName string) (*user.User, error) {
	if _, ok1 := DatabaseUser[userName]; ok1 {
		return nil, fmt.Errorf("User already exists")
	}
	user, ok2 := user.CreateUser(userName)
	if ok2 != nil {
		return nil, fmt.Errorf("User not found")
	}
	counterUsers++
	DatabaseUser[userName] = *user
	return user, nil
}

func UpdateUser(oldName string, newName string) (*user.User, error) {
	user, ok1 := DatabaseUser[oldName]
	user, ok2 := DatabaseUser[newName]
	if !ok1 || ok2 {
		return nil, fmt.Errorf("User not found or cannot change to the new name")
	}
	user.UserName = newName
	DatabaseUser[newName] = user
	delete(DatabaseUser, oldName)
	return &user, nil
}

func DeleteUser(UserName string) error {
	_, ok := DatabaseUser[UserName]
	if !ok {
		return fmt.Errorf("User not found")
	}
	delete(DatabaseUser, UserName)
	return nil
}

func removeUserTweets(tweets []twitter.Tweet, UserName string) []twitter.Tweet {
	newTweets := make([]twitter.Tweet, 0)
	for _, tweet := range tweets {
		if tweet.UserName != UserName {
			newTweets = append(newTweets, tweet)
		}
	}
	return newTweets
}

func Follow(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found")
	}
	follower.Following = append(follower.Following, followingName)
	following.Followers = append(following.Followers, followerName)
	following.NumberOfFollowers += 1
	return nil
}

func Unfollow(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found")
	}
	follower.Following = helper.RemoveString(follower.Following, followingName)
	following.Followers = helper.RemoveString(following.Followers, followerName)
	if following.NumberOfFollowers > 0 {
		following.NumberOfFollowers -= 1
	}
	return nil
}

func PostTweet(UserName string, tweetContent string) error {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return fmt.Errorf("User not found")
	}
	userTweet := twitter.Tweet{UserName: UserName, Content: tweetContent, Date: time.Now()}
	user.UserTweets = append(user.UserTweets, userTweet)
	handlePostTweet(UserName, userTweet)
	return nil
}

func handlePostTweet(UserName string, tweetContent twitter.Tweet) {
	for _, user := range DatabaseUser {
		if helper.Contains(user.Following, UserName) {
			user.Feed = append(user.Feed, tweetContent)
		}
	}
}
