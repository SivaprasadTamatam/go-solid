package main

import "errors"

// AuthenticationService handles user authentication operations
type AuthenticationService struct {
	userRepository *UserRepository
}

// NewAuthenticationService initialize a new AuthenticationService
func NewAuthenticationService(userRepository UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: &userRepository,
	}
}

// AuthenticateUser authenticate a user based on provided credentials
func (as *AuthenticationService) AuthenticateUser(username, password string) (bool, error) {
	user, err := as.getUserByUsername(username)
	if err != nil {
		return false, err
	}

	return user.Password == password, nil
}

// getUserByUsername retrieves a user from the repository
func (as *AuthenticationService) getUserByUsername(username string) (User, error) {
	for _, user := range as.userRepository.users {
		if user.UserName == username {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
