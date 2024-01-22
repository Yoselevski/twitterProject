package db

import (
	"fmt"
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
	//handleDeleteRequest(UserName)
	delete(DatabaseUser, UserName)
	return nil
}

// func handleDeleteRequest(UserName string) error {
// 	_, ok := DatabaseUser[UserName]
// 	if !ok {
// 		return fmt.Errorf("User not found")
// 	}

// 	for _ , user := range DatabaseUser {
// 		if user.UserName != UserName {
// 			newFeed := removeUserTweets(user.Feed, UserName)
// 			user.Feed := newFeed
// 		}
// 	}

//		return nil
//	}
func removeUserTweets(tweets []twitter.Tweet, UserName string) []twitter.Tweet {
	newTweets := make([]twitter.Tweet, 0)
	for _, tweet := range tweets {
		if tweet.UserName != UserName {
			newTweets = append(newTweets, tweet)
		}
	}
	return newTweets
}
