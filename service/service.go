package service

import (
	"context"

	"github.com/Pizhlo/wb-L0/models"
	"github.com/google/uuid"
)

type Service struct {
	Storage Storage
}

type Storage interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (models.Order, error)
}

func New(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}
