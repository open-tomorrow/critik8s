package config

import "os"

type AmqpConfig struct {
	Service  string
	Port     string
	Username string
	Password string
}

type RoutingKeys struct {
	DataCollector string
}

type MonitorConfig struct {
	RabbitmqPath            string `json:"rabbitmqPath"`
	RabbitmqUsername        string `json:"rabbitmqUsername"`
	RabbitmqPassword        string `json:"rabbitmqPassword"`
	DataCollectorRoutingKey string `json:"dataCollectorRoutingKey"`
}

func GetAmqpConfig() AmqpConfig {
	c := AmqpConfig{
		Service:  os.Getenv("AMQP_HOST"),
		Port:     os.Getenv("AMQP_PORT"),
		Username: os.Getenv("AMQP_USERNAME"),
		Password: os.Getenv("AMQP_PASSWORD"),
	}

	return c
}

func GetRoutingKeys() RoutingKeys {

	rk := RoutingKeys{
		DataCollector: os.Getenv("ROUTE_4"),
	}

	return rk
}

func GetMonitorConfig() MonitorConfig {

	c := MonitorConfig{
		RabbitmqPath:            os.Getenv("AMQP_PROXY_PATH"),
		RabbitmqUsername:        os.Getenv("AMQP_USERNAME"),
		RabbitmqPassword:        os.Getenv("AMQP_PASSWORD"),
		DataCollectorRoutingKey: os.Getenv("ROUTE_1"),
	}

	return c
}
