package externalServices

import (
	"Notification/entities"
	"fmt"
)

type EmailService struct {
}

func (email *EmailService) SendMessage(order *entities.Order) {
	fmt.Printf("Email sent:%v\n", order)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
