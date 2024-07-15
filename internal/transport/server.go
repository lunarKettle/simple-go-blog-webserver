package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple-go-blog-webserver/internal/database"
	"simple-go-blog-webserver/internal/models"
	"strconv"
	"strings"
	"time"
)

type Handler = func(http.ResponseWriter, *http.Request) error

func ErrorHandling(handler Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			if errors.Is(err, database.ErrEmailIsOccupied) {
				http.Error(w, err.Error(), http.StatusConflict)
				log.Println(err)
				return
			}
			if errors.Is(err, database.ErrUsernameIsOccupied) {
				http.Error(w, err.Error(), http.StatusConflict)
				log.Println(err)
				return
			}
			if errors.Is(err, database.ErrFailToGetUsers) {
				http.Error(w, "Failed to get users from database", http.StatusInternalServerError)
				log.Println("Failed to get users from database")
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
	})
}

type HTTPServer struct {
	Address string
}

func NewServer(address string) *HTTPServer {
	return &HTTPServer{Address: address}
}

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()
	eh := ErrorHandling
	mux.Handle("POST /users", eh(createUser))
	mux.Handle("GET /users", eh(getUsers))
	mux.Handle("GET /users/{id}", eh(getUserById))
	mux.Handle("POST /posts", eh(addPost))
	return http.ListenAndServe(s.Address, mux)
}

func createUser(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	userName := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	newUser := models.User{Name: name, Username: userName, Email: email}

	err := database.CreateUser(newUser)
	if err != nil {
		err := fmt.Errorf("failed to add user to database: %w", err)
		return err
	}
	//
	fmt.Fprintf(w, "User added successfully")
	return err
}

func getUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := database.GetUsers()
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		err = fmt.Errorf("failed to encode users to JSON: %w", err)
		return err
	}
	return err
}

func getUserById(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		err = fmt.Errorf("failed to get id from URL: %w", err)
		return err
	}

	user, err := database.GetUserById(id)
	if err != nil {
		err = fmt.Errorf("failed to get user from database: %w", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		err = fmt.Errorf("failed to encode users to JSON: %w", err)
		return err
	}
	return err
}

func addPost(w http.ResponseWriter, r *http.Request) error {
	text := r.URL.Query().Get("text")
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))

	if err != nil {
		err = fmt.Errorf("failed to get userId from URL: %w", err)
		return err
	}

	newPost := models.Post{Text: text, UserId: userId, Date: time.Now(), IsChanged: false}

	err = database.AddPost(newPost)
	if err != nil {
		err = fmt.Errorf("failed to add post to database: %w", err)
		return err
	}
	return err
}
