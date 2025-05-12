package post

import (
	"encoding/json"
	"net/http"

	"mini-social-network-api/internal/middleware"
	"mini-social-network-api/pkg/logger"
	"mini-social-network-api/pkg/sanitize"
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
		logger.Log.WithError(err).Error("invalid create post input")
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("content", input.Content).Info("create post request received")

	input.Content = sanitize.Sanitize(input.Content)

	userIDValue := r.Context().Value(middleware.ContextUserIDKey)
	userID, ok := userIDValue.(int64)
	if !ok {
		logger.Log.Error("user ID not found in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.svc.CreatePost(input, userID); err != nil {
		logger.Log.WithError(err).Error("create post service failed")
		http.Error(w, "could not create post", http.StatusInternalServerError)
		return
	}

	logger.Log.WithField("content", input.Content).Info("post created successfully")
	w.WriteHeader(http.StatusCreated)
}
