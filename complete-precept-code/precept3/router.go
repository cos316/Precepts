package main

import (
	"fmt"
	"net/http"
)

func simpleHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Triggered simpleHandleFunc")
	fmt.Printf("Request url: %v\n", req.URL)
}
func configure_routes() {
	// for each route pattern, register the handler
	http.HandleFunc("/routeA", simpleHandleFunc)
}
