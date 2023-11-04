package services

import (
	"Notification/entities"
	"Notification/externalServices"
	"fmt"
)

type OrderService struct {
	EmailService *externalServices.EmailService
	SMSService   *externalServices.SMSService
}

func (o *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Order created: %v\n", order)
	// send SMS or Email
	o.EmailService.SendMessage(order)
	o.SMSService.SendMessage(order)

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{
		externalServices.NewEmailService(),
		externalServices.NewSMSService(),
	}
}
