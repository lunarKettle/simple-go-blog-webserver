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
	server.Start()
}
