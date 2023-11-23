package main

// InMemoryUserRepository handles the user management operations

type InMemoryUserRepository struct {
	users []User
}

// NewInMemoryUserRepository initializes a new UserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make([]User, 0)}
}

// AddUser adds a new user to the repository

func (ur *InMemoryUserRepository) AddUser(user User) {
	ur.users = append(ur.users, user)
}
