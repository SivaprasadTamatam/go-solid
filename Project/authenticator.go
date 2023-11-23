package main

// Authenticator defines the interface for different authentication methods
type Authenticator interface {
	Authenticate(user User, credentials interface{}) bool
}
