package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-go-blog-webserver/internal/models"
	"simple-go-blog-webserver/internal/repository"
	"strconv"
	"strings"
	"time"
)

type HTTPServer struct {
	Address        string
	database       repository.Database
	userRepository repository.UserRepository
	postRepository repository.PostRepository
}

func NewServer(address string) *HTTPServer {
	return &HTTPServer{Address: address}
}

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()
	eh := errorHandling
	mux.Handle("POST /users", eh(s.createUser))
	mux.Handle("GET /users", eh(s.getUsers))
	mux.Handle("GET /users/{id}", eh(s.getUserById))
	mux.Handle("POST /posts", eh(s.addPost))
	mux.Handle("GET /posts/{id}", eh(s.getPostById))
	s.database.OpenConnection()
	s.userRepository = repository.NewUserRepository(s.database)
	s.postRepository = repository.NewPostRepository(s.database)
	defer s.database.CloseConnection()
	return http.ListenAndServe(s.Address, mux)
}

func (s *HTTPServer) createUser(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	userName := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	newUser := models.User{Name: name, Username: userName, Email: email}

	err := s.userRepository.CreateUser(newUser)
	if err != nil {
		err := fmt.Errorf("failed to add user to database: %w", err)
		return err
	}
	//
	fmt.Fprintf(w, "User added successfully")
	return err
}

func (s *HTTPServer) getUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.userRepository.GetUsers()
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

func (s *HTTPServer) getUserById(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		err = fmt.Errorf("failed to get id from URL: %w", err)
		return err
	}

	user, err := s.userRepository.GetUserById(id)
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

func (s *HTTPServer) addPost(w http.ResponseWriter, r *http.Request) error {
	text := r.URL.Query().Get("text")
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))

	if err != nil {
		err = fmt.Errorf("failed to get userId from URL: %w", err)
		return err
	}

	newPost := models.Post{Text: text, UserId: userId, Date: time.Now(), IsChanged: false}

	err = s.postRepository.AddPost(newPost)
	if err != nil {
		err = fmt.Errorf("failed to add post to database: %w", err)
		return err
	}
	return err
}

func (s *HTTPServer) getPostById(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		err = fmt.Errorf("failed to get id from URL: %w", err)
		return err
	}

	user, err := s.postRepository.GetPostById(id)
	if err != nil {
		err = fmt.Errorf("failed to get post from database: %w", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		err = fmt.Errorf("failed to encode posts to JSON: %w", err)
		return err
	}
	return err
}
