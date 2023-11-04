package main

import (
	"Notification/entities"
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

	orderService := services.NewOrderService()
	orderService.CreateOrder(order)
}
