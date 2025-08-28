package broker

import "encoding/json"

type MessageType int

const (
	GameStartMessage MessageType = iota
)

type Message struct {
	Type MessageType
}

// Parses a message from JSON data
func ParseMessage(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
