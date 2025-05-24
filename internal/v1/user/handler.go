package user

import (
	"encoding/json"
	"mini-social-network-api/internal/httphelper"
	"mini-social-network-api/pkg/logger"
	"mini-social-network-api/pkg/sanitize"
	"mini-social-network-api/pkg/validate"
	"net/http"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.Log.WithError(err).Error("invalid register input")
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	if err := validate.Validator.Struct(input); err != nil {
		logger.Log.WithError(err).Error("validation failed for register input")
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, "validation failed")
		return
	}

	logger.Log.WithField("username", input.Username).Info("register request received")

	input.Username = sanitize.Sanitize(input.Username)
	input.Email = sanitize.Sanitize(input.Email)

	if err := h.svc.Register(input); err != nil {
		logger.Log.WithError(err).Error("register service failed")
		httphelper.WriteErrorResponse(w, http.StatusInternalServerError, "register failed")
		return
	}

	logger.Log.WithField("username", input.Username).Info("user registered successfully")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logger.Log.WithError(err).Error("invalid login input")
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, "invalid input")
		return
	}

	if err := validate.Validator.Struct(input); err != nil {
		logger.Log.WithError(err).Error("validation failed for login input")
		httphelper.WriteErrorResponse(w, http.StatusBadRequest, "validation failed")
		return
	}

	logger.Log.WithField("email", input.Email).Info("login request received")

	user, token, err := h.svc.Login(input)
	if err != nil {
		logger.Log.WithError(err).Error("unauthrorized login attempt")
		httphelper.WriteErrorResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	logger.Log.WithField("user_id", user.ID).Info("user logged in successfully")

	resp := map[string]interface{}{
		"user": map[string]interface{}{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		},
		"token": token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
