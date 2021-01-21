package main

import "service1"

func main() {
	// Start Producer
	go service1.StartProducer()

	go service1.StartConsumer()

	// Start API server
	service1.Start()
}
