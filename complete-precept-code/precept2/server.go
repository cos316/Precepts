package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to setup a listener - %v\n", err)
	}
	defer ln.Close()
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection - %v\n", err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("Failed to read from connection - %v\n", err)
	}
	fmt.Println(string(buf))
}
