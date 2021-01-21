package service2

import (
	"encoding/json"
	"log"
	"service2/model"

	"github.com/streadway/amqp"
)

func StartConsumer() {
	conn, err := amqp.Dial(rabbitConfig.uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"request_queue", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	messageChannel, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for {
		select {
		case message := <-messageChannel:
			var messageHello model.HelloMessage
			if err := json.Unmarshal(message.Body, &messageHello); err != nil {
				panic(err)
			}
			log.Printf("INFO: received msg: %+v", messageHello)
			err = message.Ack(true)
			failOnError(err, "Fail to ack")

			// Process data
			messageHello.Message = "Hello " + messageHello.Message

			// Send back
			responseChan <- messageHello
		}

	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
