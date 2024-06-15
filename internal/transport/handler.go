package transport

import (
	"fmt"
	"net/http"
)

type appHandler struct {
	// Можно добавить поля для управления состоянием приложения
}

func (ah *appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ah.handleGet(w, r)
	case http.MethodPost:
		ah.handlePost(w, r)
	case http.MethodPut:
		ah.handlePut(w, r)
	case http.MethodDelete:
		ah.handleDelete(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (ah *appHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handling GET request")
}

func (ah *appHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handling POST request")
}

func (ah *appHandler) handlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handling PUT request")
}

func (ah *appHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handling DELETE request")
}
