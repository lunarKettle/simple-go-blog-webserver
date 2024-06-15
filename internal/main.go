package main

import "go-todo-server/internal/transport"

func main() {
	server := transport.NewServer(":8080")
	server.Start()
}
