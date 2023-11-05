package externalServices

import (
	"fmt"
)

type EmailService struct {
}

func (email *EmailService) SendNotification(receiver string, message string) {
	fmt.Printf("Email sent to:%s message:%s\n", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
