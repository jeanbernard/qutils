package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const SensorListQueue = "SensorList"
const SensorDiscoveryExchange = "SensorDiscovery"

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establish connection to the message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get channel for connection")

	fmt.Println("Connection succesful to RabbitMQ")
	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,       // Name of the queue
		false,      // durable: Queues will survive server restarts
		autoDelete, // autoDelete: delete queues if no active listeners
		false,      // exclusive: Only conn to work with this queue
		false,      // noWait: Send msg without waiting for queue to be set up. It not rdy it will fail.
		nil,        //amqp.Table: additional configs for the exchange
	)
	failOnError(err, "Failed to declare queue")

	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	} else {
		fmt.Println("Connection to RabbitMQ successfully")
	}
}
