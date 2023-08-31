package publisher

import (
	"encoding/json"

	"github.com/Pizhlo/wb-L0/internal/app/stan/data"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/nats-io/stan.go"
)

type Publish interface {
	SendMsg() error
}

type Publisher struct {
	stan.Conn
}

type Msg struct {
	Order models.Order
}

func New(nc stan.Conn) (*Publisher, error) {
	return &Publisher{nc}, nil
}

func (p *Publisher) SendMsg() error {
	var msg Msg
	order := data.RandomOrder()
	msg.Order = order

	data, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	return p.Publish("test", data)
}
