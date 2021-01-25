package broker

import "context"

type Broker interface {
	Publish(ctx context.Context, topic string, message []byte) error
	Subscribe(topic string, handler EventHandler) error
}

type EventHandler func([]byte)
