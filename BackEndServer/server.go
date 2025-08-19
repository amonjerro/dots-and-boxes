package main

import (
	"log"
	"os"

	"github.com/amonjerro/dots-and-boxes/internal/broker"
)

func main() {
	// Pull up the messaging broker and configure it
	brokerURI := os.Getenv("BROKER_URI")
	consume_exchange := os.Getenv("CONSUME_EXCHANGE")
	log.Printf("%s", brokerURI)
	consumer, err := broker.NewConsumer(brokerURI, "client_messages", consume_exchange)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err.Error())
	}
	consumer.AddHandler(broker.GameStartMessage, broker.GameStartHandler{})
	consumer.Consume()
}
