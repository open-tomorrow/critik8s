package amqp

import (
	"context"
	"data-collector/config"
	e "data-collector/error"
	k8s "data-collector/kubernetes"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishSubscribeTest() {
	ch, err := conn.Channel()
	e.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	e.FailOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	nodes := k8s.GetNodes()

	routingKeys := config.GetRoutingKeys()

	for {
		err = ch.PublishWithContext(ctx,
			"logs_topic",        // exchange
			routingKeys.Monitor, // routing key for monitor
			false,               // mandatory
			false,               // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(nodes),
			})
		e.FailOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent nodes")
		time.Sleep(5 * time.Second)
	}
}
