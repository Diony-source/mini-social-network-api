package follow

import (
	"encoding/json"
	"net/http"

	"mini-social-network-api/internal/middleware"
	"mini-social-network-api/pkg/logger"
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
		logger.Log.WithError(err).Error("invalid follow input")
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	logger.Log.WithField("followee_id", input.FolloweeID).Info("follow request received")

	userIDValue := r.Context().Value(middleware.ContextUserIDKey)
	followerID, ok := userIDValue.(int64)
	if !ok {
		logger.Log.Error("user ID not found in context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	if err := h.svc.FollowUser(followerID, input); err != nil {
		logger.Log.WithError(err).Error("failed to follow user")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.WithField("follower_id", followerID).Info("user followed successfully")
	w.WriteHeader(http.StatusCreated)
}
