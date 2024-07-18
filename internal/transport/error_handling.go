package transport

import (
	"errors"
	"log"
	"net/http"
	"simple-go-blog-webserver/internal/repository"
)

type Handler = func(http.ResponseWriter, *http.Request) error

func errorHandling(handler Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			if errors.Is(err, repository.ErrEmailIsOccupied) {
				http.Error(w, err.Error(), http.StatusConflict)
				log.Println(err)
				return
			}
			if errors.Is(err, repository.ErrUsernameIsOccupied) {
				http.Error(w, err.Error(), http.StatusConflict)
				log.Println(err)
				return
			}
			if errors.Is(err, repository.ErrFailToGetUsers) {
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
