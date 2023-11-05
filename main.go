package main

import (
	"Notification/entities"
	"Notification/externalServices"
	"Notification/services"
)

func main() {
	order := &entities.Order{
		"21",
		"Dev",
		"WWW@test.com",
		656,
		true,
	}

	orderService := services.NewOrderService(externalServices.NewSMSService())
	orderService.CreateOrder(order)
}
