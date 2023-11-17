package main

import "fmt"

// EmailNotificationService implements the NotificationService for sending notifications via email
type EmailNotificationService struct{}

func (e *EmailNotificationService) SendNotification(message string) error {
	fmt.Println("Email Notification :", message)
	return nil
}
