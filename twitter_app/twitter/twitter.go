package twitter

import (
	"time"
	"sort"
)

type Tweet struct {
	UserName string
	Content string
	Date    time.Time
}

func RemoveTweet(tweets []Tweet, tweet Tweet) []Tweet {
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

func SortTweetsByDate(tweets []Tweet) []Tweet {
	sortedTweets := make([]Tweet, len(tweets))
	copy(sortedTweets, tweets)

	sort.Slice(sortedTweets, func(i, j int) bool {
		return sortedTweets[i].Date.Before(sortedTweets[j].Date)
	})

	return sortedTweets
}
