package main

// AuthenticationManager defines the interface for authentication operations
type AuthenticationManager interface {
	AddAuthenticator(authenticator Authenticator)
	AuthenticateUser(user User, credentials interface{}) bool
}
