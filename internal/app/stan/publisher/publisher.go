package publisher

import (
	"encoding/json"
	"time"

	"github.com/Pizhlo/wb-L0/internal/app/stan/data"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
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

func New(nc stan.Conn) *Publisher {
	return &Publisher{nc}
}

func (p *Publisher) Start(ticker *time.Ticker, done chan bool) error {
	go func() error {
		for {
			select {
			case <-done:
				return nil
			case <-ticker.C:
				err := p.sendMsg()
				if err != nil {
					return err
				}
			}
		}
	}()

	return nil
}

func (p *Publisher) sendMsg() error {
	var msg Msg
	order := data.RandomOrder()
	msg.Order = order

	//fmt.Printf("sending random order: %+v\n", order)

	data, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "publisher: cannot marshal msg to json")
	}

	return p.Publish("test", data)
}
