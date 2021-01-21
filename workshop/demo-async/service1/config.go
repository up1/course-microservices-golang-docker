package service1

import "os"

type RabbitConfig struct {
	uri string
	requestQueue string
	responseQueue string
}

var rabbitConfig = RabbitConfig{
	uri:          getEnv("RABBIT_URI", "amqp://guest:guest@localhost:5672/"),
	requestQueue: getEnv("RABBIT_REQUEST_QUEUE", "request_queue"),
	responseQueue: getEnv("RABBIT_RESPONSE_QUEUE", "response_queue"),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
