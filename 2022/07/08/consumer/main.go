package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"group.id":           "files-client",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": true,
	}

	client, err := kafka.NewConsumer(&config)
	if err != nil {
		log.Fatalln("Kafka.NewConsumer", err)
	}

	if err := client.Subscribe("files", nil); err != nil {
		log.Fatalln("client.Subscribe", err)
	}

	// XXX:
	// Refer to https://youtu.be/jr7OULxYm0A for a full example
	// using a Kafka Consumer

	run := true
	for run {
		ev := client.Poll(150)

		switch msg := ev.(type) {
		case *kafka.Message:
			fmt.Println("Message", string(msg.Value))
			run = false
		case kafka.Error:
			fmt.Println("Error", msg)
			run = false
		}
	}

	if err := client.Close(); err != nil {
		log.Fatalln("Close", err)
	}
}
