package main

// User struct that represents a user in the system

type User struct {
	UserName string
	Password string
}

// UserRepository handles user management operations

type UserRepository struct {
	users []User
}

// NewUserRepository initialization of new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]User, 0),
	}
}

// AddUser adds a new user to the repository
func (ur *UserRepository) AddUser(user User) {
	ur.users = append(ur.users, user)
}
