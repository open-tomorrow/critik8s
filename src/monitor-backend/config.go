package main

import "os"

type Config struct {
	RabbitmqPath     string `json:"rabbitmqPath"`
	RabbitmqUsername string `json:"rabbitmqUsername"`
	RabbitmqPassword string `json:"rabbitmqPassword"`
}

func config() Config {

	c := Config{
		RabbitmqPath:     os.Getenv("rmq_relay"),
		RabbitmqUsername: os.Getenv("rmq_username"),
		RabbitmqPassword: os.Getenv("rmq_password"),
	}

	return c
}
