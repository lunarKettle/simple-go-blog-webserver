package transport

import (
	"net/http"
)

type HTTPServer struct {
	Address string
}

func NewServer(address string) *HTTPServer {
	return &HTTPServer{Address: address}
}

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("GET /users/{id}", getUserById)
	mux.HandleFunc("POST /posts", addPost)
	return http.ListenAndServe(s.Address, mux)
}
