package main

import "fmt"

// SMSService implements the NotificationService for sending notifications via SMS
type SMSService struct{}

func (s *SMSService) SendNotification(message string) error {
	fmt.Println("SMS Notification :", message)
	return nil
}
