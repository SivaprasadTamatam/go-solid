package main

import (
	"fmt"
)

func main() {
	// Create a new user repository
	userRepository := NewInMemoryUserRepository()

	user := User{
		Username: "Siva",
		Password: "secure_password",
	}

	userRepository.AddUser(user)

	// Create a new authentication service with the user repository
	authService := NewAuthenticationService(userRepository)
	// Password based authenticator
	authService.AddAuthenticator(&PasswordAuthenticator{})
	// Fingerprint based authenticator
	authService.AddAuthenticator(&FingerprintAuthenticator{})

	isAuthenticatedPassword := authService.AuthenticateUser(user, "secure_password")

	if isAuthenticatedPassword {
		fmt.Println("Password Authentication Successful, Welcome,", user.Username)
	} else {
		fmt.Println("Password Authentication failed, Invalid credentials")
	}

	isAuthenticatedFingerprint := authService.AuthenticateUser(user, "fingerprint_data")

	if isAuthenticatedFingerprint {
		fmt.Println("Fingerprint Authentication Successful, Welcome,", user.Username)
	} else {
		fmt.Println("Fingerprint Authentication failed, Invalid credentials")
	}

	user1 := User{
		Username: "Prasad",
		Password: "another_secure_password",
	}
	userManager := userRepository
	userManager.AddUser(user1)

	isAuthenticatedPassword = authService.AuthenticateUser(user1, "another_secure_password")

	if isAuthenticatedPassword {
		fmt.Println("Password Authentication Successful, Welcome,", user1.Username)
	} else {
		fmt.Println("Password Authentication failed, Invalid credentials")
	}

}
