package user

import (
	"errors"
	"mini-social-network-api/pkg/auth"
	"mini-social-network-api/pkg/logger"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Register(input RegisterRequest) error {
	hashed, err := auth.HashPassword(input.Password)
	if err != nil {
		logger.Log.WithError(err).Error("failed to hash password")
		return err
	}

	user := &User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hashed,
	}

	logger.Log.WithField("email", user.Email).Info("attempting to create user")

	if err := s.repo.CreateUser(user); err != nil {
		logger.Log.WithError(err).Error("failed to create user in DB")
		return err
	}

	logger.Log.WithField("email", user.Email).Info("user created successfully")
	return nil
}

func (s *Service) Login(input LoginRequest) (*User, string, error) {
	user, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		logger.Log.WithField("email", input.Email).Warn("user not found during login")
		return nil, "", errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(input.Password, user.PasswordHash) {
		logger.Log.WithField("email", input.Email).Warn("invalid password attempt")
		return nil, "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		logger.Log.WithError(err).Error("token generation failed")
		return nil, "", err
	}

	return user, token, nil
}
