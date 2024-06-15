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
	handler := &appHandler{}
	fmt.Printf("Starting server at %s\n", s.Address)
	return http.ListenAndServe(s.Address, handler)
}
