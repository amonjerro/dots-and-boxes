package broker

type MockConsumer struct {
	handlers    map[MessageType]MessageHandler
	ConsumeFunc func()
}

func (m *MockConsumer) AddHandler(mt MessageType, mh MessageHandler) {
	m.handlers[mt] = mh
}

func (m *MockConsumer) Consume() {
	m.ConsumeFunc()
}
