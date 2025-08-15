package broker

import amqp "github.com/rabbitmq/amqp091-go"

type MockConnection struct {
	ChannelFunc func() (Channel, error)
	CloseFunc   func() error
}

type MockChannel struct {
	ExchangeDeclareFunc func(string, string, bool, bool, bool, bool, amqp.Table) error
	QueueDeclareFunc    func(string, bool, bool, bool, bool, amqp.Table) (amqp.Queue, error)
	QueueBindFunc       func(string, string, string, bool, amqp.Table) error
	ConsumeFunc         func(string, string, bool, bool, bool, bool, amqp.Table) (<-chan amqp.Delivery, error)
	CloseFunc           func() error
}

func (m *MockConnection) Channel() (Channel, error) {
	return m.ChannelFunc()
}

func (m *MockConnection) Close() error {
	return m.CloseFunc()
}

func (m *MockChannel) ExchangeDeclare(s1 string, s2 string, b1 bool, b2 bool, b3 bool, b4 bool, t amqp.Table) error {
	return m.ExchangeDeclareFunc(s1, s2, b1, b2, b3, b4, t)
}

func (m *MockChannel) QueueDeclare(s1 string, b1 bool, b2 bool, b3 bool, b4 bool, t amqp.Table) (amqp.Queue, error) {
	return m.QueueDeclareFunc(s1, b1, b2, b3, b4, t)
}

func (m *MockChannel) QueueBind(s1 string, s2 string, s3 string, b1 bool, t amqp.Table) error {
	return m.QueueBindFunc(s1, s2, s3, b1, t)
}

func (m *MockChannel) Consume(s1 string, s2 string, b1 bool, b2 bool, b3 bool, b4 bool, t amqp.Table) (<-chan amqp.Delivery, error) {
	return m.ConsumeFunc(s1, s2, b1, b2, b3, b4, t)
}

func (m *MockChannel) Close() error {
	return m.CloseFunc()
}
