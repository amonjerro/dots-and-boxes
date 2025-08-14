package broker

type Consumer interface {
	AddHandler(MessageType, MessageHandler)
	Consume()
}

type ServerConsumer struct {
	conn         Connection
	ch           Channel
	queueName    string
	exchangeName string
	handlers     map[MessageType]MessageHandler
}

func NewConsumer(brokerURI string, queue string, exchange string) (Consumer, error) {
	conn, err := Connect(brokerURI)
	failOnError(err, "Failed to connect to broker")
	channel, err := conn.Channel()
	failOnError(err, "Failed to connect to channel")
	SetupExchange(channel, exchange)

	q, err := channel.QueueDeclare(
		queue,
		true,  // durable
		false, // delete when unusued
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	failOnError(err, "Failed to create consumer queue")

	err = channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil,
	)
	failOnError(err, "Failed to bind queue")

	return &ServerConsumer{
		conn:         conn,
		ch:           channel,
		queueName:    q.Name,
		exchangeName: exchange,
	}, nil
}

func (c *ServerConsumer) AddHandler(key MessageType, handler MessageHandler) {
	c.handlers[key] = handler
}

func (c *ServerConsumer) Consume() {
	msgs, err := c.ch.Consume(
		c.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume from channel")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			msg, err := ParseMessage(d.Body)
			if err != nil {

			}
			handler, ok := c.handlers[msg.Type]
			if ok {
				handler.Handle(msg)
			}
		}
	}()

	<-forever
}
