package main

func main() {
	// Existing code using EmailNotificationService
	emailNotification := &EmailNotificationService{}
	notificationSender := &NotificationSender{notificationService: emailNotification}
	notificationSender.SendNotification("Hello, this is an email notification")

	// Extending to use SMSService without modifying the existing code
	smsNotificationService := &SMSService{}
	notificationSender.notificationService = smsNotificationService
	notificationSender.SendNotification("Hello, this is SMS notification")
}
