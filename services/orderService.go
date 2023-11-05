package services

import (
	"Notification/entities"
	"Notification/externalServices"
)

type OrderService struct {
	Notifier externalServices.Notifier
}

func (o *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	o.Notifier.SendNotification(order.UserName, "order created successfully")

	return order
}

func NewOrderService(notifier externalServices.Notifier) *OrderService {
	return &OrderService{
		Notifier: notifier,
	}
}
