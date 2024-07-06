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
	return http.ListenAndServe(s.Address, mux)
}
