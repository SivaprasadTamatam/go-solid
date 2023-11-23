package main

// UserRepository handles user management operations
type UserRepository interface {
	AddUser(user User)
}
