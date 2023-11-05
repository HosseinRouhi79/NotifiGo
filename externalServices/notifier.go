package externalServices

type Notifier interface {
	SendNotification(receiver string, message string)
}
