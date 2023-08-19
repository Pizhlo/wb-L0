package service

import (
	"context"

	"github.com/Pizhlo/wb-L0/models"
)

type Service struct {
	Storage Storage
}

type Storage interface {
	GetUserByID(ctx context.Context, id int) (models.User, error)
}

func New(storage Storage) *Service {
	return &Service{
		Storage: storage,
	}
}
