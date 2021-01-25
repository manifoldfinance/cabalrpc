package broker

import (
	"context"
	"github.com/nats-io/nats.go"
)

type natsBroker struct {
	*nats.Conn
}

func NewNatsBroker(url string) (Broker, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return natsBroker{
		Conn: conn,
	}, nil
}

func (n natsBroker) Publish(_ context.Context, topic string, message []byte) error {
	return n.Conn.Publish(topic, message)
}

func (n natsBroker) Subscribe(topic string, handler EventHandler) error {
	_, err := n.Conn.Subscribe(topic, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	return err
}
