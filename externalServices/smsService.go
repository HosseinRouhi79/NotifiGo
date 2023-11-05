package externalServices

import (
	"fmt"
)

type SMSService struct {
}

func (sms *SMSService) SendNotification(receiver string, message string) {
	fmt.Printf("SMS sent to:%s message:%s\n", receiver, message)
}

func NewSMSService() *SMSService {
	return &SMSService{}
}
