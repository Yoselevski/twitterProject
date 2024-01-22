package user

import (
	"fmt"
	"sort"
	"time"
	"twitter_app/helper"
	"twitter_app/helper/twitter"
)

type User struct {
	UserName          string
	Followers         []string
	numberOfFollowers int
	Following         []string
	userTweets        []twitter.Tweet
	Feed              []twitter.Tweet
}

func CreateUser(UserName string) (*User, error) {
	user := &User{
		UserName:          UserName,
		Followers:         []string{},
		numberOfFollowers: 0,
		Following:         []string{},
		userTweets:        []twitter.Tweet{},
		feed:              []twitter.Tweet{},
	}
	return user, nil
}

func Follow(followerName string, followingName string) error {
	follower, ok1 := DatabaseUser[followerName]
	following, ok2 := DatabaseUser[followingName]
	if !ok1 || !ok2 {
		return fmt.Errorf("User not found")
	}
	follower.Following = append(follower.Following, followingName)
	following.Followers = append(following.Followers, followerName)
	following.numberOfFollowers += 1
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
	if following.numberOfFollowers > 0 {
		following.numberOfFollowers -= 1
	}
	return nil
}

func getFollowerNames(UserName string) ([]string, error) {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return nil, fmt.Errorf("User not found")
	}
	return user.Followers, nil
}

func getFollowingNames(UserName string) ([]string, error) {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return nil, fmt.Errorf("User not found")
	}
	return user.Following, nil
}

func PostTweet(UserName string, tweetContent string) error {
	user, ok := DatabaseUser[UserName]
	if !ok {
		return fmt.Errorf("User not found")
	}
	userTweet := Tweet{UserName: UserName, Content: tweetContent, Date: time.Now()}
	user.userTweets = append(user.userTweets, userTweet)
	handlePostTweet(UserName, userTweet)
	return nil
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

	follower.feed = append(follower.feed, following.userTweets...)
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
		return sortedUsers[i].numberOfFollowers > sortedUsers[j].numberOfFollowers
	})

	// Return the top 'n' users
	if n > len(sortedUsers) {
		return sortedUsers, nil
	}
	return sortedUsers[:n], nil
}
