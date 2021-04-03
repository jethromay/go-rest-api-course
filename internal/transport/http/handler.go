package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - Stores pointer for comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns pointer to handler
func NewHandler() *Handler {
	return &Handler{}
}

// Sets up all of the routes for the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
