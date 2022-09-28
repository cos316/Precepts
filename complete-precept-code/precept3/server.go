package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// get port number from command line
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./simple [server port]")
	}
	server_port := os.Args[1]
	fmt.Println("Setting up server to listen on", server_port)
	configure_routes()
	err := http.ListenAndServe("localhost:"+server_port, nil)
	if err != nil {
		log.Fatal("Failed to setup http server")
	}
}
