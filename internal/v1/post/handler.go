package post

import (
	"encoding/json"
	"net/http"

	"mini-social-network-api/internal/middleware"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var input CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	userIDValue := r.Context().Value(middleware.ContextUserIDKey)
	userID, ok := userIDValue.(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.svc.CreatePost(input, userID); err != nil {
		http.Error(w, "could not create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
