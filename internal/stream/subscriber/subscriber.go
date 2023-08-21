package subscriber

import (
	"context"

	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	"github.com/Pizhlo/wb-L0/internal/stream/publisher"
	"github.com/nats-io/nats.go"
)

type Subscribe interface {
	ReceiveMsg(m *publisher.Msg, storage postgres.Storage) error
}

type Subscriber struct {
	*nats.EncodedConn
}

func New(c *nats.EncodedConn, storage postgres.Storage) *Subscriber {
	subscriber := &Subscriber{c}

	c.Subscribe("foo", func(msg *publisher.Msg) {
		subscriber.ReceiveMsg(msg, storage)
	})

	return subscriber

}

func (s *Subscriber) ReceiveMsg(m *publisher.Msg, storage postgres.Storage) error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	err := storage.SaveOrder(ctx, m.Order)
	if err != nil {
		return err
	}

	return nil
}

