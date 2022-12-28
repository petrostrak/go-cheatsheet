# RabbitMQ in Golang
`RabbitMQ` is one of the most popular open-source, lightweight message brokers which helps in transferring data based on a messaging protocol. `RabbitMQ` supports multiple messaging protocols, and one of the most commonly used underlying protocol to route messages is `Advanced Messaging Queueing Protocol (AMQP)`.

`AMQP` has three primary entities – namely: `Queues`, `Exchanges` and `Bindings`. `AMQP` being a programmable protocol, provides flexibility to configure these entities to help transfer messages from producer to consumer.

When a publisher publishes a message to `RabbitMQ`, it is first received at an exchange. The exchange routes the message to variously connected queues which is then received by the consumers.

## Set up a RabbitMQ instance locally
To get started, we’ll first need to get a `RabbitMQ` instance running on our machine that we can interact and play with it. The easiest and quickest approach is to use the docker run command and specifying the image name that we want to run.
```
docker run -d --hostname rabbit-in-go --name a-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management 
```
This gets our `RabbitMQ` instance locally with its management interface available at `http://localhost:15672`.

## Publish a message

To interact with the instance, we would now need to open a channel over the connection created. This channel allows us to publish messages to an exchange and also declare queues to which we can subscribe to.

For this we will use the **default exchange**. The default exchange is a pre-declared direct exchange with no name, referred by an empty string. When using the default exchange, the message is delivered to the queue with a name equal to the routing key.

*Note: Every queue is automatically bound to the default exchange with a routing key same as the queue name.*

For creating a new queue, we use the packages’s `QueueDeclare` function defined in the channel with our desired queue properties. Once the queue is created, we can now publish a message over the channel using its `Publish` function.

## Consume messages from a queue

For consuming messages, we use the packages `Consume` function defined in the channel with our desired queue properties. Then we can iterate over the messages and read them.

## To run localy
To publish a message:
```
cd producer && go run producer.go
```
To consume a message:
```
cd consumer && go run consumer.go
```