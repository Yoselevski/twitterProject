package userCommands

import (
	"fmt"
	"sort"
	"time"
	"twitter_app/helper"
	"twitter_app/twitter"
	"twitter_app/user"
)

var DatabaseUser = make(map[string]user.User)

func InsertUser(userName string) (*user.User, error) {
	if _, ok1 := DatabaseUser[userName]; ok1 {
		return nil, fmt.Errorf("User already exists InsertUser")
	}
	user, ok2 := user.CreateUser(userName)
	if ok2 != nil {
		return nil, fmt.Errorf("User not found InsertUser")
	}
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
	DatabaseUser[newName] = user // by value
	delete(DatabaseUser, oldName)
	return &user, nil
}

func DeleteUser(UserName string) error {
	_, ok := DatabaseUser[UserName]
	if !ok {
		return fmt.Errorf("User not found DeleteUser")
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
	if followerName == followingName {
		return fmt.Errorf("User cannot follow himself")
	}
	handleFollowRequest(followerName, followingName)
	return nil
}

func Unfollow(followerName string, followingName string) error {
	if followerName == followingName {
		return fmt.Errorf("User cannot unfollow himself")
	}
	handleUnFollowRequest(followerName, followingName)
	return nil
}

func PostTweet(UserName string, tweetContent string) error {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return fmt.Errorf("User not found PostTweet")
	}
	userTweet := twitter.Tweet{UserName: UserName, Content: tweetContent, Date: time.Now()}
	user.UserTweets = append(user.UserTweets, userTweet)
	DatabaseUser[UserName] = user
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

func handleFollowRequest(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found handleFollowRequest")
	}
	follower.Following = append(follower.Following, followingName)
	newFeed := append(follower.Feed, following.UserTweets...)
	follower.Feed = twitter.SortTweetsByDate(newFeed)
	DatabaseUser[followerName] = follower
	following.Followers = append(following.Followers, followerName)
	following.NumberOfFollowers += 1
	DatabaseUser[followingName] = following
	return nil
}

func handleUnFollowRequest(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found handleUnFollowRequest")
	}
	newFeed := follower.Feed
	for _, tweet := range following.Feed {
		newFeed = twitter.RemoveTweet(follower.Feed, tweet)
	}
	follower.Following = helper.RemoveString(follower.Following, followingName)
	follower.Feed = newFeed
	DatabaseUser[followerName] = follower
	following.Followers = helper.RemoveString(following.Followers, followerName)
	if following.NumberOfFollowers > 0 {
		following.NumberOfFollowers -= 1
	}
	DatabaseUser[followingName] = following
	return nil
}

func GetUserFeed(UserName string) ([]twitter.Tweet, error) {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return nil, fmt.Errorf("User not found GetUserFeed")
	}
	return user.Feed, nil
}

func GetMutualFollowers(UserName1 string, UserName2 string) ([]string, error) {
	user1, ok1 := DatabaseUser[UserName1]
	user2, ok2 := DatabaseUser[UserName2]
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("User not found GetMutualFollowers")
	}

	mutualFollowers := []string{}
	for _, user := range DatabaseUser {
		if helper.Contains(user1.Following, user.UserName) && helper.Contains(user2.Following, user.UserName) {
			mutualFollowers = append(mutualFollowers, user.UserName)
		}
	}

	if len(mutualFollowers) == 0 {
		return nil, fmt.Errorf("No mutual followers mutualFollowers")
	}

	return mutualFollowers, nil
}

func GetTopInfluencers(n int) ([]user.User, error) {
	// Sort users by number of followers in descending order
	sortedUsers := make([]user.User, 0, len(DatabaseUser))
	for _, user := range DatabaseUser {
		sortedUsers = append(sortedUsers, user)
	}
	sort.Slice(sortedUsers, func(i, j int) bool {
		return sortedUsers[i].NumberOfFollowers > sortedUsers[j].NumberOfFollowers
	})

	// Return the top 'n' users
	if n > len(sortedUsers) {
		return sortedUsers, nil
	}
	return sortedUsers[:n], nil
}
