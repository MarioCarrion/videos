package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := fmt.Sprintf("%s:%d", "google.com", 80) // WRONG: Use net.JoinHostPort
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to", conn.RemoteAddr())
}
