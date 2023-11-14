package amqp

import (
	"context"
	"fmt"
	"monitor-backend/config"
	e "monitor-backend/error"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var routingKeys = config.GetRoutingKeys()
var conn *amqp.Connection

func Init() *amqp.Connection {
	conn = connect()
	return conn
}

func connect() *amqp.Connection {
	c := config.GetAmqpConfig()

	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", c.Username, c.Password, c.Service, c.Port)

	conn, err := amqp.Dial(address)
	e.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func RpcRequest(entity string, id string) string {
	ch, err := conn.Channel()
	e.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	e.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	e.FailOnError(err, "Failed to register a consumer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",                        // exchange
		routingKeys.DataCollector, // routing key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: entity,
			ReplyTo:       q.Name,
			Body:          []byte(id),
		})
	e.FailOnError(err, "Failed to publish a message")

	delivery := <-msgs

	return string(delivery.Body)
}
