package main

import (
	"context"
	amqpClient "data-collector/amqp"
	"data-collector/config"
	e "data-collector/error"
	k8s "data-collector/kubernetes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn := amqpClient.Connect()
	defer conn.Close()

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

	body, err := json.Marshal(nodes)
	if err != nil {
		fmt.Println(err)
		return
	}

	routingKeys := config.GetRoutingKeys()

	for {
		err = ch.PublishWithContext(ctx,
			"logs_topic",        // exchange
			routingKeys.Monitor, // routing key for monitor
			false,               // mandatory
			false,               // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(string(body)),
			})
		e.FailOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent nodes")
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
