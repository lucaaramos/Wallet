package rabbit

import (
	"context"
	"encoding/json"
	"log"
)

func StartConsumers(ctx context.Context) error {
	ch := GetChannel()

	for _, c := range GetConsumers() {

		q, err := ch.QueueDeclare(
			c.RoutingKey,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		err = ch.QueueBind(
			q.Name,
			c.RoutingKey,
			"wallet.events",
			false,
			nil,
		)
		if err != nil {
			return err
		}

		msgs, err := ch.Consume(
			q.Name,
			"",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		go func(handler HandlerFunc, routingKey string) {
			log.Println("👂 listening:", routingKey)

			for msg := range msgs {

				var event Event
				if err := json.Unmarshal(msg.Body, &event); err != nil {
					log.Println("invalid event:", err)
					msg.Nack(false, false)
					continue
				}

				dataBytes, _ := json.Marshal(event.Data)

				err := handler(ctx, dataBytes)

				if err != nil {
					log.Println("handler error:", err)
					msg.Nack(false, true)
				} else {
					msg.Ack(false)
				}
			}
		}(c.Handler, c.RoutingKey)
	}

	return nil
}
