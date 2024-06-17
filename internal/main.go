package main

import "simple-go-blog-webserver/internal/transport"

func main() {
	server := transport.NewServer(":8080")
	server.Start()
}
