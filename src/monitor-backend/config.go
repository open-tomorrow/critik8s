package main

import "os"

type Config struct {
	RabbitmqPath            string `json:"rabbitmqPath"`
	RabbitmqUsername        string `json:"rabbitmqUsername"`
	RabbitmqPassword        string `json:"rabbitmqPassword"`
	DataCollectorRoutingKey string `json:"dataCollectorRoutingKey"` // TO BE REMOVED, ONLY FOR TESTS
}

func getConfig() Config {

	c := Config{
		RabbitmqPath:            os.Getenv("AMQ_PROXY_PATH"),
		RabbitmqUsername:        os.Getenv("RMQ_USERNAME"),
		RabbitmqPassword:        os.Getenv("RMQ_PASSWORD"),
		DataCollectorRoutingKey: os.Getenv("ROUTE_1"),
	}

	return c
}
