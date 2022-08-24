package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	var ip net.IP

	// Type "net.IPv4" implements the interface "encoding.TextMarshaler",
	// specifically the method "MarshalText"
	flag.TextVar(&ip, "ip", net.IPv4(192, 168, 0, 100), "`IP address` to parse")
	flag.Parse()

	fmt.Printf("Value: %v\n", ip)
}
