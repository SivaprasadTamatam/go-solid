package main

// NotificationService defines the contract for sending notifications

type NotificationService interface {
	SendNotification(message string) error
}
