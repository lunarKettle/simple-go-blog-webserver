package main

import (
	"fmt"
	"simple-go-blog-webserver/internal/database"
	"simple-go-blog-webserver/internal/transport"
)

func main() {
	server := transport.NewServer(":8080")
	err := database.OpenConnection()
	if err != nil {
		fmt.Println("Ошибка базы данных")
	}
	defer database.CloseConnection()
	fmt.Printf("Starting server at %s\n", server.Address)
	server.Start()
}
