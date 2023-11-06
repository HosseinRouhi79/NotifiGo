package entities

import "Notification/enums"

type Order struct {
	ID               string
	UserName         string
	UserEmail        string
	Price            float64
	Status           bool
	NotificationType enums.NotificationType
}
