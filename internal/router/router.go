package router

import (
	"net/http"

	"crud-app/internal/handlers"
)

// New builds the HTTP router with all CRUD routes wired to their handlers.
func New() http.Handler {
	mux := http.NewServeMux()

	itemHandler := handlers.NewItemHandler()

	mux.HandleFunc("POST /items", itemHandler.Create)
	mux.HandleFunc("GET /items", itemHandler.List)
	mux.HandleFunc("GET /items/{id}", itemHandler.Get)
	mux.HandleFunc("PUT /items/{id}", itemHandler.Update)
	mux.HandleFunc("DELETE /items/{id}", itemHandler.Delete)

	return mux
}
