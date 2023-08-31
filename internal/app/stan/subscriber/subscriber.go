package subscriber

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Pizhlo/wb-L0/internal/app/stan/publisher"
	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	"github.com/Pizhlo/wb-L0/internal/service"
	"github.com/nats-io/stan.go"
)

type Subscribe interface {
	ReceiveMsg(m *publisher.Msg, storage postgres.Repo) error
}

type Subscriber struct {
	stan.Conn
	service service.Service
}

func New(nc stan.Conn, service service.Service) error {
	subscriber := &Subscriber{nc, service}

	_, err := subscriber.Subscribe("test", func(m *stan.Msg) {
		fmt.Println("recieved msg")
		err := subscriber.ReceiveMsg(m.Data)
		if err != nil {
			fmt.Println("unable to receive msg: ", err)
		}
	})
	if err != nil {
		return err
	}

	return nil

}

func (s *Subscriber) ReceiveMsg(m []byte) error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var msg publisher.Msg

	err := json.Unmarshal(m, &msg)
	if err != nil {
		return err
	}

	err = s.service.SaveOrder(ctx, msg.Order)
	if err != nil {
		return err
	}

	return nil
}
