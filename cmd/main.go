package main

import (
	"fmt"
	"log"
)

func main() {
	userService := InitializeUserService()

	user, err := userService.CreateUser("John Doe", "john@example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created user: %+v\n", user)

	// Try to get the user
	foundUser, err := userService.GetUser(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found user: %+v\n", foundUser)
}
