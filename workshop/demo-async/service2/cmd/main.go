package main

import "service2"

func main() {
	go service2.StartProducer()
	service2.StartConsumer()
}
