package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mini-social-network-api/internal/middleware"
	"mini-social-network-api/pkg/logger"
	"mini-social-network-api/pkg/sanitize"
	"mini-social-network-api/pkg/validate"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
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

	if err := validate.Validator.Struct(input); err != nil {
		logger.Log.WithError(err).Error("validation failed for create post input")
		http.Error(w, "validation failed", http.StatusBadRequest)
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

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("🔍 UpdatePost handler reached")

	postIDStr := chi.URLParam(r, "id")
	logger.Log.WithFields(logrus.Fields{
		"path_param_id": postIDStr,
		"url":           r.URL.String(),
		"path":          r.URL.Path,
	}).Info("DEBUG: received post id from URL")

	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		logger.Log.WithError(err).Error("invalid post ID")
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userID := ctx.Value(middleware.ContextUserIDKey).(int64)
	role := ctx.Value(middleware.ContextUserRoleKey).(string)

	post, err := h.svc.GetByID(postID)
	if err != nil {
		logger.Log.WithError(err).Error("failed to get post")
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	if post.AuthorID != userID && role != "admin" {
		logger.Log.Error("user not authorized to update post")
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var input UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.Log.WithError(err).Error("invalid update post input")
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := validate.Validator.Struct(input); err != nil {
		logger.Log.WithError(err).Error("validation failed for update post input")
		http.Error(w, "validation failed", http.StatusBadRequest)
		return
	}

	input.Content = sanitize.Sanitize(input.Content)
	post.Content = input.Content
	if err := h.svc.UpdatePost(input, postID); err != nil {
		logger.Log.WithError(err).Error("failed to update post")
		http.Error(w, "failed to update post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "id")
	logger.Log.WithField("path_param_id", postIDStr).Info("delete post request received")

	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		logger.Log.WithError(err).Error("invalid post ID")
		http.Error(w, "invalid post ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userID := ctx.Value(middleware.ContextUserIDKey).(int64)
	role := ctx.Value(middleware.ContextUserRoleKey).(string)

	post, err := h.svc.GetByID(postID)
	if err != nil {
		logger.Log.WithError(err).Error("failed to get post")
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	if post.AuthorID != userID && role != "admin" {
		logger.Log.Error("user not authorized to delete post")
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	if err := h.svc.DeletePost(postID); err != nil {
		logger.Log.WithError(err).Error("failed to delete post")
		http.Error(w, "failed to delete post", http.StatusInternalServerError)
		return
	}

	logger.Log.WithField("post_id", postID).Info("post deleted successfully")
	w.WriteHeader(http.StatusNoContent)
}
