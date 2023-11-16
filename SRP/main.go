package main

import (
	"fmt"
)

func main() {
	// create a user repository
	userRepository := NewUserRepository()

	user := User{
		UserName: "Sivaprasad",
		Password: "secure_password",
	}

	userRepository.AddUser(user)

	// Create a authentication service for the user repository
	authService := NewAuthenticationService(*userRepository)

	// Attempt to authenticate user
	username := "Siva"
	password := "secure_password"

	isAuthenticated, err := authService.AuthenticateUser(username, password)

	if err != nil {
		fmt.Println("Authentication Failed : ", err.Error())
		return
	}
	if isAuthenticated {
		fmt.Println("Authentication Successful. Welcome, ", username)
	} else {
		fmt.Println("Authentication Failed. Invalid credentials")
	}

}
