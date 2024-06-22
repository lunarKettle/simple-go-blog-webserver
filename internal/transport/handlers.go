package transport

import (
	"fmt"
	"log"
	"net/http"
	"simple-go-blog-webserver/internal/database"
	"simple-go-blog-webserver/internal/models"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	userName := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	newUser := models.User{Name: name, UserName: userName, Email: email}
	fmt.Println("User added successfully")

	err := database.CreateUser(newUser)
	if err != nil {
		log.Fatal("Failed to add user to database")
	} else {
		fmt.Fprintf(w, "name: %v\n", newUser.Name)
		fmt.Fprintf(w, "username: %v\n", newUser.UserName)
		fmt.Fprintf(w, "email: %v\n", newUser.Email)
	}
}
