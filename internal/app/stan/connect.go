package stream

import (
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type NatsConnection struct {
	*nats.Conn
}

func Connect() (*NatsConnection, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return &NatsConnection{}, errors.Wrap(err, "unable to connect nats")
	}

	ncStruct := &NatsConnection{nc}

	return ncStruct, nil
}

func (nc *NatsConnection) Close() {
	nc.Conn.Close()
}
