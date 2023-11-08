package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	service := os.Getenv("RABBITMQ_CLUSTER_SERVICE_HOST")
	port := os.Getenv("RABBITMQ_CLUSTER_SERVICE_PORT")

	username := os.Getenv("RMQ_USERNAME")
	password := os.Getenv("RMQ_PASSWORD")
	monitorRoutingKey := os.Getenv("ROUTE_1")

	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, service, port)

	conn, err := amqp.Dial(address)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
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
	failOnError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "ciao"

	for {
		err = ch.PublishWithContext(ctx,
			"logs_topic",      // exchange
			monitorRoutingKey, // routing key for monitor, TO BE REMOVED, ONLY FOR TESTS
			false,             // mandatory
			false,             // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent %s", body)
		time.Sleep(5 * time.Second)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "ciao"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}
