package externalServices

import (
	"Notification/entities"
	"fmt"
)

type SMSService struct {
}

func (sms *SMSService) SendMessage(order *entities.Order) {
	fmt.Printf("SMS sent :%v\n", order)
}

func NewSMSService() *SMSService {
	return &SMSService{}
}
