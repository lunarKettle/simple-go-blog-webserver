package transport

import (
	"encoding/json"
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
		http.Error(w, "Failed to add users to database", http.StatusInternalServerError)
		log.Println("Failed to add user to database", err)
		return
	}
	fmt.Fprintf(w, "name: %v\n", newUser.Name)
	fmt.Fprintf(w, "username: %v\n", newUser.UserName)
	fmt.Fprintf(w, "email: %v\n", newUser.Email)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := database.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get users from database", http.StatusInternalServerError)
		log.Println("Failed to get users from database")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
		log.Println("Failed to encode users to JSON:", err)
		return
	}
}
