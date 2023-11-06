package externalServices

import "Notification/enums"

type Notifier interface {
	SendNotification(receiver string, message string)
}

func NewNotifier(notificationType enums.NotificationType) Notifier {
	switch notificationType {
	case enums.Email:
		return NewEmailService()
	case enums.SMS:
		return NewSMSService()
	default:
		return nil

	}
}
