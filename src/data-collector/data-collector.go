package main

import (
	"data-collector/amqp"
)

func main() {
	var forever chan struct{}

	conn := amqp.Init()
	defer conn.Close()

	go amqp.RpcAux()

	// TODO remove, for proves Publish Subscribe only
	amqp.PublishSubscribeTest()

	<-forever
}
