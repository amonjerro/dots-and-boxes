package broker

type MessageHandler interface {
	Handle(m *Message) error
}

type GameStartHandler struct{}

func (handler GameStartHandler) Handle(msg *Message) error {
	return nil
}
