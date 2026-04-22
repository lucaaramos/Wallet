package rabbit

import amqp "github.com/rabbitmq/amqp091-go"

var conn *amqp.Connection
var ch *amqp.Channel

func Init(uri string) error {
	var err error

	conn, err = amqp.Dial(uri)
	if err != nil {
		return err
	}

	ch, err = conn.Channel()
	if err != nil {
		return err
	}

	return nil
}

func GetChannel() *amqp.Channel {
	return ch
}

func Close() {
	if ch != nil {
		ch.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
