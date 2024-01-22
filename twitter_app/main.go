package main

import (
	"fmt"
	"twitter_app/userCommands"
)

func main() {
	fmt.Println("Hey there! I hope you enjoy to tweet your thoughts and follow each other!")

	// Create users
	_, err1 := userCommands.InsertUser("user1")
	if err1 != nil {
		fmt.Println("Error creating user1:", err1)
	}

	_, err2 := userCommands.InsertUser("user2")
	if err2 != nil {
		fmt.Println("Error creating user2:", err2)
		return
	}

	// Update user
	_, err3 := userCommands.UpdateUser("user1", "updatedUser1")
	if err3 != nil {
		fmt.Println("Error updating user:", err3)
	}

	// Delete user
	err4 := userCommands.DeleteUser("user2")
	if err4 != nil {
		fmt.Println("Error deleting user:", err4)
	}

	// Post tweet
	err5 := userCommands.PostTweet("updatedUser1", "Tweeting1 user1!")
	if err5 != nil {
		fmt.Println("Error posting tweet:", err5)
	}

	// Post tweet
	err6 := userCommands.PostTweet("updatedUser1", "Tweeting2 user1!")
	if err6 != nil {
		fmt.Println("Error posting tweet:", err6)
	}

	// Follow user that doesnt exist
	err7 := userCommands.Follow("updatedUser1", "user3")
	if err7 != nil {
		fmt.Println("Error following user:", err7)
	}
	// Create user again
	_, err8 := userCommands.InsertUser("user2")
	if err8 != nil {
		fmt.Println("Error creating user1:", err8)
	}

	// Follow updatedUser1
	err9 := userCommands.Follow("user2", "updatedUser1")
	if err9 != nil {
		fmt.Println("Error following user:", err9)
	}

	_, err10 := userCommands.GetUserFeed("user2")
	if err10 != nil {
		fmt.Println("Error following user:", err10)
	}

	// Unfollow user
	err11 := userCommands.Unfollow("user2", "updatedUser1")
	if err11 != nil {
		fmt.Println("Error unfollowing user:", err11)
	}

	_, err12 := userCommands.GetUserFeed("user2")
	if err12 != nil {
		fmt.Println("Error following user:", err12)
	}
	// Get top influencers
	_, err13 := userCommands.GetTopInfluencers(2)
	if err13 != nil {
		fmt.Println("Error getting top influencers:", err13)
	}
	fmt.Println("Users created, tweet posted, and followed each other successfully!")

}
