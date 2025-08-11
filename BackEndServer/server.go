package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s %s", msg, err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	// Pull up the messaging broker and configure it
	brokerURI := os.Getenv("BROKER_URI")
	conn, err := amqp.Dial(brokerURI)
	failOnError(err, "Failed to connect to messaging service")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel with messaging service")

	defer ch.Close()
	defer conn.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HttpHandler(resp http.ResponseWriter, r *http.Request) {
	fmt.Fprint(resp, "Hello World!")
}
