package amqp

import (
	"data-collector/config"
	e "data-collector/error"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect() *amqp.Connection {
	c := config.GetAmqpConfig()

	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", c.Username, c.Password, c.Service, c.Port)

	conn, err := amqp.Dial(address)
	e.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}
