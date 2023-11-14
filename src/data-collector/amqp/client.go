package amqp

import (
	"context"
	"data-collector/config"
	e "data-collector/error"
	k8s "data-collector/kubernetes"
	"fmt"
	"log"
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

func RpcAux() {
	ch, err := conn.Channel()
	e.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		routingKeys.MonitorBackend, // name
		false,                      // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	e.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	e.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	e.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {

			response := k8s.GetResource(d.CorrelationId)

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
				})
			e.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
