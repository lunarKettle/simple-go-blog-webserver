package transport

import (
	"fmt"
	"log"
	"net/http"
	"simple-go-blog-webserver/internal/database"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	userName := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	fmt.Println("User added successfully")
	fmt.Fprintf(w, "name: %v\n", name)
	fmt.Fprintf(w, "username: %v\n", userName)
	fmt.Fprintf(w, "email: %v\n", email)
	err := database.CreateUser(name, userName, email)
	if err != nil {
		log.Fatal("Failed to add user to database")
	}
}
