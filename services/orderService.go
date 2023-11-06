package services

import (
	"Notification/entities"
	"Notification/externalServices"
)

type OrderService struct {
	externalServices.Notifier
}

func (o *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	o.Notifier = externalServices.NewNotifier(order.NotificationType)
	o.SendNotification(order.UserName, "order created successfully")

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
