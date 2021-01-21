package service2

import (
	"encoding/json"
	"log"
	"service2/model"

	"github.com/streadway/amqp"
)

// channel to publish rabbit messages
var responseChan = make(chan model.HelloMessage, 10)

func StartProducer() {
	conn, err := amqp.Dial(rabbitConfig.uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"response_queue", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for {
		select {
		case message := <-responseChan:
			helloMessage, _ := json.Marshal(message)
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         helloMessage,
				})
			failOnError(err, "Failed to publish a message")
			log.Printf("Sent %s back to service1", message)
		}
	}

}
