package transport

import (
	"net/http"
	"simple-go-blog-webserver/internal/repository"
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
	mux.Handle("GET /posts", eh(s.getPostByUserId))
	mux.Handle("GET /posts/{id}", eh(s.getPostById))
	timedMux := timingMiddleware(mux)

	s.database.OpenConnection()
	s.userRepository = repository.NewUserRepository(s.database)
	s.postRepository = repository.NewPostRepository(s.database)
	defer s.database.CloseConnection()
	return http.ListenAndServe(s.Address, timedMux)
}
