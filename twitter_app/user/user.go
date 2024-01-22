package user

import (
	"fmt"
	"sort"
	"time"
	"twitter_app/helper"
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



func handlePostTweet(UserName string, tweetContent Tweet) {
	for _, user := range DatabaseUser {
		if helper.Contains(user.Following, UserName) {
			user.feed = append(user.feed, tweetContent)
		}
	}
}

func handleFollowRequest(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found")
	}

	follower.feed = append(follower.feed, following.UserTweets...)
	sortTweetsByDate(follower.feed)
	return nil
}
func handleUnFollowRequest(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found")
	}
	newFeed := follower.feed
	for _, tweet := range following.feed {
		newFeed = removeTweet(follower.feed, tweet)
	}
	follower.feed = newFeed
	return nil
}

func removeTweet(tweets []Tweet, tweet Tweet) []Tweet {
	index := -1
	for i, t := range tweets {
		if t == tweet {
			index = i
			break
		}
	}
	if index != -1 {
		tweets = append(tweets[:index], tweets[index+1:]...)
	}
	return tweets
}

func sortTweetsByDate(tweets []Tweet) {
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].Date.Before(tweets[j].Date)
	})
}

func GetUserFeed(UserName string) ([]Tweet, error) {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return nil, fmt.Errorf("User not found")
	}
	return user.feed, nil
}

func GetMutualFollowers(UserName1 string, UserName2 string) ([]*User, error) {
	user1, ok1 := DatabaseUser[UserName1]
	user2, ok2 := DatabaseUser[UserName2]
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("User not found")
	}

	mutualFollowers := []*User{}
	for _, user := range DatabaseUser {
		if helper.Contains(user1.Following, user.UserName) && helper.Contains(user2.Following, user.UserName) {
			mutualFollowers = append(mutualFollowers, &user)
		}
	}

	if len(mutualFollowers) == 0 {
		return nil, fmt.Errorf("No mutual followers")
	}

	return mutualFollowers, nil
}

func GetTopInfluencers(n int) ([]User, error) {
	// Sort users by number of followers in descending order
	sortedUsers := make([]User, 0, len(DatabaseUser))
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
