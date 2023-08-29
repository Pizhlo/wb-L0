package publisher

import (
	"github.com/Pizhlo/wb-L0/internal/stream/data"
	"github.com/Pizhlo/wb-L0/models"
	"github.com/nats-io/nats.go"
)

type Publish interface {
	SendMsg() error
}

type Publisher struct {
	*nats.EncodedConn
}

type Msg struct {
	Order models.Order
}

func New(nc *nats.Conn) (*Publisher, error) {
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return &Publisher{}, err
	}
	return &Publisher{c}, nil
}

func (p *Publisher) SendMsg() error {
	var msg Msg
	order := data.RandomOrder()
	msg.Order = order

	return p.Publish("foo", msg)
}
