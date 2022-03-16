package main

import (
	"fmt"

	"net"
	"net/netip"
)

func main() {
	ip := "192.168.0.1"

	fmt.Println(net.ParseIP(ip))

	addr, _ := netip.ParseAddr(ip)
	fmt.Println(addr)
}
