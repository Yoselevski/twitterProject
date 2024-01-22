package main

import (
	"fmt"
	"twitter_app/user"
	"twitter_app/db"

)


func main() {
	fmt.Println("Users created, tweet posted, and followed each other successfully!")

	// Create users
	_, err := db.InsertUser("user1")
	if err != nil {
		fmt.Println("Error creating user1:", err)
		return
	}

	_, err := db.InsertUser("user2")
	if err != nil {
		fmt.Println("Error creating user2:", err)
		return
	}

	fmt.Println(db.DatabaseUser)
}



