package broker

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go.elastic.co/apm"
)

func NewApmBroker(core Broker) Broker {
	return apmBroker{core: core}
}

type apmBroker struct {
	core Broker
}

func (a apmBroker) Publish(ctx context.Context, topic string, message []byte) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "publish")
	span = span.SetTag("topic", topic)
	defer span.Finish()
	return a.core.Publish(ctx, topic, message)
}

func (a apmBroker) Subscribe(topic string, handler EventHandler) error {
	span, _ := apm.StartSpan(context.Background(), "subscribe", "broker")
	defer span.End()
	return a.core.Subscribe(topic, handler)
}
