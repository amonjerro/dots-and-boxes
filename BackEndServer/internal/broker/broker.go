package broker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s %s", msg, err)
	}
}

func Connect(uri string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to messaging service")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	return conn, ch
}

/*
We want to set up two exchanges:
 1. For client games to publish actions and be consumed by this server
 2. For this server to update the state of the game to all clients
*/
func SetupExchange(ch *amqp.Channel, exchangeName string) {
	err := ch.ExchangeDeclare(
		exchangeName, // exchange name
		"topic",      // exchange type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Server publish exchange declaration failed")
}
