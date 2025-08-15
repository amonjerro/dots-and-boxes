package broker

import "testing"

func TestAddHandler(t *testing.T) {
	tests := []struct {
		mType             MessageType
		mHandler          MessageHandler
		addedHandlerTotal int
	}{
		// Add test values here
		{GameStartMessage, &GameStartHandler{}, 1},
	}
	// Arrange
	consumer := &MockConsumer{
		handlers: make(map[MessageType]MessageHandler),
	}

	for _, tt := range tests {
		consumer.AddHandler(tt.mType, tt.mHandler)
		if len(consumer.handlers) != tt.addedHandlerTotal {
			t.Errorf("Expected %d handlers added, found only %d", tt.addedHandlerTotal, len(consumer.handlers))
		}
	}
}
