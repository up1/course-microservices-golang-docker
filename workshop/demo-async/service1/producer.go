package service1

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func StartProducer() {
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

	for {
		select {
		case message := <-pchan:
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
			log.Printf("Sent %s to queue", message)
		}
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
