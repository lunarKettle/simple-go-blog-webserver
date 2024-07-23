package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-go-blog-webserver/internal/models"
	"strconv"
	"strings"
	"time"
)

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

func (s *HTTPServer) getPostByUserId(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		err = fmt.Errorf("failed to get userId from URL: %w", err)
		return err
	}

	posts, err := s.postRepository.GetPostsByUserId(userId)
	if err != nil {
		err = fmt.Errorf("failed to get posts from database: %w", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		err = fmt.Errorf("failed to encode posts to JSON: %w", err)
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

	post, err := s.postRepository.GetPostById(id)
	if err != nil {
		err = fmt.Errorf("failed to get post from database: %w", err)
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(post); err != nil {
		err = fmt.Errorf("failed to encode posts to JSON: %w", err)
		return err
	}
	return err
}
