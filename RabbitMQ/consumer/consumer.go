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

	// declare consumer with its properties over channel
	msgs, err := channel.Consume(
		"rabbitMQ in Golang!", // queue
		"",                    // consumer
		true,                  // autp ack
		false,                 // exclusive
		false,                 // no local
		false,                 // no wait
		nil,                   //args
	)
	if err != nil {
		panic(err)
	}

	// print consumed messages from queue
	done := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received message: %s\n", msg.Body)
		}
	}()

	fmt.Println("Waiting for messages...")
	<-done
}
