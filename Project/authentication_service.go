package main

// AuthenticationService handles user authentication operations
type AuthenticationService struct {
	userRepository UserRepository
	authenticators []Authenticator
}

// NewAuthenticationService initializes a AuthenticationService
func NewAuthenticationService(userRepository UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		authenticators: make([]Authenticator, 0),
	}
}

// AddAuthenticator adds a new authenticator to the AuthenticationService
func (as *AuthenticationService) AddAuthenticator(authenticator Authenticator) {
	as.authenticators = append(as.authenticators, authenticator)
}

// AuthenticateUser authenticate a user based on provided credentials
func (as *AuthenticationService) AuthenticateUser(user User, credentials interface{}) bool {
	for _, authenticator := range as.authenticators {
		if authenticator.Authenticate(user, credentials) {
			return true
		}
	}
	return false
}
