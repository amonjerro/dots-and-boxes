package broker

type MessageHandler interface {
	Handle(m *Message)
}

type GameStartHandler struct{}

func (handler GameStartHandler) Handle(msg *Message) {

}
