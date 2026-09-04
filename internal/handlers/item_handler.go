package handlers

import "net/http"

// ItemHandler groups the CRUD endpoints for the "item" resource.
type ItemHandler struct{}

func NewItemHandler() *ItemHandler {
	return &ItemHandler{}
}

// Create handles POST /items
func (h *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List handles GET /items
func (h *ItemHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get handles GET /items/{id}
func (h *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update handles PUT /items/{id}
func (h *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete handles DELETE /items/{id}
func (h *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
