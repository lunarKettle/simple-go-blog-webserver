package transport

import (
	"fmt"
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
	err := http.ListenAndServe(s.Address, mux)
	if err == nil {
		fmt.Printf("Starting server at %s\n", s.Address)
	}
	return err
}
