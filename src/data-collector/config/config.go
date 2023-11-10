package config

import "os"

type AmqpConfig struct {
	Service  string
	Port     string
	Username string
	Password string
}

type RoutingKeys struct {
	Monitor string
}

func GetAmqpConfig() AmqpConfig {

	c := AmqpConfig{
		Service:  os.Getenv("RABBITMQ_CLUSTER_SERVICE_HOST"),
		Port:     os.Getenv("RABBITMQ_CLUSTER_SERVICE_PORT"),
		Username: os.Getenv("AMQP_USERNAME"),
		Password: os.Getenv("AMQP_PASSWORD"),
	}

	return c
}

func GetRoutingKeys() RoutingKeys {

	rk := RoutingKeys{
		Monitor: os.Getenv("ROUTE_1"),
	}

	return rk
}
