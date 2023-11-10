package config

import "os"

type Config struct {
	RabbitmqPath            string `json:"rabbitmqPath"`
	RabbitmqUsername        string `json:"rabbitmqUsername"`
	RabbitmqPassword        string `json:"rabbitmqPassword"`
	DataCollectorRoutingKey string `json:"dataCollectorRoutingKey"` // TO BE REMOVED, ONLY FOR TESTS
}

func Get() Config {

	c := Config{
		RabbitmqPath:            os.Getenv("AMQP_PROXY_PATH"),
		RabbitmqUsername:        os.Getenv("AMQP_USERNAME"),
		RabbitmqPassword:        os.Getenv("AMQP_PASSWORD"),
		DataCollectorRoutingKey: os.Getenv("ROUTE_1"),
	}

	return c
}
