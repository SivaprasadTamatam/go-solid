package main

// PasswordAuthenticator is an implementation of Authenticator for password-based authentication
type PasswordAuthenticator struct{}

// Authenticate authenticates a used based on a password
func (pa *PasswordAuthenticator) Authenticate(user User, credentials interface{}) bool {
	password, ok := credentials.(string)
	if !ok {
		return false
	}
	return user.Password == password
}
