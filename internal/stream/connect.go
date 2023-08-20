package stream

import (
	"github.com/nats-io/nats.go"
)

type NatsConnection struct {
	*nats.Conn
}

func Connect() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nc, err
	}

	return nc, nil
}

func (nc *NatsConnection) Close() {
	nc.Close()
}
