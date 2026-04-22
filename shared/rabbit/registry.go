package rabbit

import "context"

type HandlerFunc func(ctx context.Context, data []byte) error

type Consumer struct {
	RoutingKey string
	Handler    HandlerFunc
}

var consumers []Consumer

func Register(c Consumer) {
	consumers = append(consumers, c)
}

func GetConsumers() []Consumer {
	return consumers
}
