package internal

import "github.com/Pizhlo/wb-L0/internal/service"

type Order struct {
	service service.Service
}

func NewOrder(service service.Service) *Order {
	return &Order{service: service}
}
