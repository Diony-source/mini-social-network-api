package follow

import (
	"encoding/json"
	"net/http"

	"mini-social-network-api/internal/httphelper"
	"mini-social-network-api/internal/middleware"
	"mini-social-network-api/pkg/logger"
	"mini-social-network-api/pkg/validate"
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
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	if err := validate.Validator.Struct(input); err != nil {
		logger.Log.WithError(err).Error("validation failed for follow input")
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, "validation failed")
		return
	}

	logger.Log.WithField("followee_id", input.FolloweeID).Info("follow request received")

	userIDValue := r.Context().Value(middleware.ContextUserIDKey)
	followerID, ok := userIDValue.(int64)
	if !ok {
		logger.Log.Error("user ID not found in context")
		httphelper.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.svc.FollowUser(followerID, input); err != nil {
		logger.Log.WithError(err).Error("failed to follow user")
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	logger.Log.WithField("follower_id", followerID).Info("user followed successfully")
	httphelper.WriteJSONResponse(w, http.StatusCreated, map[string]string{"status": "followed"})
}
