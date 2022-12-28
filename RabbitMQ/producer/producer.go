package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ in Golang!")

	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ instance.")

	// open a channel over the connection established to interact with RabbitMQ
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// declare queue with its properties over the channel
	queue, err := channel.QueueDeclare(
		"rabbitMQ in Golang!", // name
		false,                 // durable
		false,                 // auto delete
		false,                 // exclusive
		false,                 // no wait
		nil,                   // args
	)
	if err != nil {
		panic(err)
	}

	// publish a message
	err = channel.Publish(
		"",                    //exchange
		"rabbitMQ in Golang!", // key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Just a default exchange with RabbitMQ in golang!"),
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
}
