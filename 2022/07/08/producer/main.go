package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	config := kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	}

	client, err := kafka.NewProducer(&config)
	if err != nil {
		log.Fatalln("Kafka.NewProducer", err)
	}

	var topic = "files"

	if err := client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: newMessage(),
	}, nil); err != nil {
		log.Fatalln("producer.Producer", err)
	}

	fmt.Println("message sent, exiting")

	client.Close()
}

func newMessage() []byte {
	return []byte("https://youtu.be/jr7OULxYm0A")

	// file, err := os.Open("Golang_Microservices_Events_Streaming_using_Apache_Kafka.mp4")
	// if err != nil {
	// 	log.Fatalln("os.Open", err)
	// }
	// defer file.Close()

	// data, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Fatalln("ioutil.ReadAll", err)
	// }

	// return data
}
