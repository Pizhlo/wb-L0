package service

import (
	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	"github.com/Pizhlo/wb-L0/internal/stream/publisher"
	"github.com/Pizhlo/wb-L0/internal/stream/subscriber"
)

type Service struct {
	Storage    postgres.Storage
	Publisher  publisher.Publish
	Subscriber subscriber.Subscribe
}

func New(storage postgres.Storage, publisher publisher.Publish, subscriber subscriber.Subscribe) *Service {
	return &Service{
		Storage:    storage,
		Publisher:  publisher,
		Subscriber: subscriber,
	}
}
