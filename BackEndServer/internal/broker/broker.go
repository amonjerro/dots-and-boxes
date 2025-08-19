package broker

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection interface {
	Channel() (Channel, error)
	Close() error
}

type Channel interface {
	ExchangeDeclare(string, string, bool, bool, bool, bool, amqp.Table) error
	QueueDeclare(string, bool, bool, bool, bool, amqp.Table) (amqp.Queue, error)
	QueueBind(string, string, string, bool, amqp.Table) error
	Consume(string, string, bool, bool, bool, bool, amqp.Table) (<-chan amqp.Delivery, error)
	Close() error
}

type amqpConnection struct {
	*amqp.Connection
}

type amqpChannel struct {
	*amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s %s", msg, err)
	}
}

func (c *amqpConnection) Channel() (Channel, error) {
	ch, err := c.Connection.Channel()
	if err != nil {
		return nil, err
	}
	return &amqpChannel{ch}, nil
}

func Connect(uri string) (Connection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return &amqpConnection{conn}, nil
}

func SetupExchange(ch Channel, exchangeName string) {
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
