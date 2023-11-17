package main

// NotificationSender utilizes the NotificationService for sending notifications
type NotificationSender struct {
	notificationService NotificationService
}

func (ns *NotificationSender) SendNotification(message string) error {
	return ns.notificationService.SendNotification(message)
}
