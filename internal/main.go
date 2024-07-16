package main

import (
	"fmt"
	"simple-go-blog-webserver/internal/transport"
)

func main() {
	server := transport.NewServer(":8080")
	fmt.Printf("Starting server at %s\n", server.Address)
	server.Start()
}
