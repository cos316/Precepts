package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err  := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server - %v\n", err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "Hello world!")
}
