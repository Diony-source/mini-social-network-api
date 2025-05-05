package follow

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

func (h *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	var input FollowRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	userIDValue := r.Context().Value(middleware.ContextUserIDKey)
	followerID, ok := userIDValue.(int64)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.svc.FollowUser(followerID, input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
