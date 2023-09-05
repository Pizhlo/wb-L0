package subscriber

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Pizhlo/wb-L0/internal/app/stan/publisher"
	"github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
)

type Subscribe interface {
	ReceiveMsg(m *publisher.Msg, storage postgres.Repo) error
}

type orderSvc interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error)
	SaveOrder(ctx context.Context, order models.Order) error
	Recover(ctx context.Context) error
	GetAllOrders(ctx context.Context) ([]models.Order, error)
}

type Subscriber struct {
	stan.Conn
	service orderSvc
}

func New(nc stan.Conn, service orderSvc) error {
	subscriber := &Subscriber{nc, service}

	_, err := subscriber.Subscribe("test", func(m *stan.Msg) {
		fmt.Println("recieved msg")
		err := subscriber.receiveMsg(m.Data)
		if err != nil {
			fmt.Println("unable to receive msg: ", err)
		}
	})
	if err != nil {
		return errors.Wrap(err, "Subscriber: cannot subscribe")
	}

	return nil

}

func (s *Subscriber) receiveMsg(m []byte) error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var msg publisher.Msg

	err := json.Unmarshal(m, &msg)
	if err != nil {
		return errors.Wrap(err, "Subscriber: cannot unmarshal json")
	}

	err = s.service.SaveOrder(ctx, msg.Order)
	if err != nil {
		return errors.Wrap(err, "Subscriber: cannot save order")
	}

	return nil
}
