package main

// FingerprintAuthenticator is an implementation of Authenticator for fingerprint-based authentication
type FingerprintAuthenticator struct{}

// Authenticator authenticates a user based on fingerprint
func (fa *FingerprintAuthenticator) Authenticate(user User, credentials interface{}) bool {
	return true
}
