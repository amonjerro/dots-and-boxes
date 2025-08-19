package broker

import (
	"testing"
)

func TestConnection(t *testing.T) {
	// Arrange
	mockCh := &MockChannel{
		CloseFunc: func() error {
			return nil // No error when closing channel
		},
	}
	mockConnection := &MockConnection{
		ChannelFunc: func() (Channel, error) {
			return mockCh, nil // Successfully return mock channel
		},
		CloseFunc: func() error {
			return nil // No errors when closing
		},
	}

	// Act
	ch, err := mockConnection.Channel()

	// Assess
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if err := ch.Close(); err != nil {
		t.Fatalf("Unexpected error closing channel: %v", err)
	}
}
