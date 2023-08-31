package stream

import (
	"github.com/nats-io/nats.go"
)

type NatsConnection struct {
	*nats.Conn
}

func Connect() (*NatsConnection, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return &NatsConnection{}, err
	}

	ncStruct := &NatsConnection{nc}

	return ncStruct, nil
}

func (nc *NatsConnection) Close() {
	nc.Conn.Close()
}
