package rabbit

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

var exchange = "wallet.events"

func Publish(ctx context.Context, routingKey string, data interface{}) error {
	ch := GetChannel()

	event := Event{
		EventID: uuid.New().String(),
		Type:    routingKey,
		Data:    data,
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
