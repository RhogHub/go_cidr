package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <CIDR>")
		os.Exit(1)
	}

	cidr := os.Args[1]
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("Error parsing CIDR:", err)
		os.Exit(1)
	}

	network := ipnet.IP
	broadcast := make(net.IP, len(network))
	for i := range network {
		broadcast[i] = network[i] | ^ipnet.Mask[i]
	}

	ones, bits := ipnet.Mask.Size()

	fmt.Println("CIDR:", cidr)
	fmt.Println("Network Address:", network)
	fmt.Println("Broadcast Address:", broadcast)
	fmt.Println("Number of Hosts:", (1<<uint32(bits-ones))-2)
	fmt.Println("Subnet Mask:", ipnet.Mask)
}
