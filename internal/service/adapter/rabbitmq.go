package adapter

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

const RabbitMQAdapterName = "rabbitmq"

type rabbitMQAdapter struct {
	exchange   string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func MakeRabbitMQAdapter(url, exchange string) (Adapter, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQAdapter{
		exchange:   exchange,
		connection: conn,
		channel:    ch,
	}, nil
}

func (adap *rabbitMQAdapter) Send(event interface{}) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = adap.channel.Publish(
		adap.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (adap *rabbitMQAdapter) Stop() error {
	if err := adap.channel.Close(); err != nil {
		return err
	}

	if err := adap.connection.Close(); err != nil {
		return err
	}

	return nil
}
