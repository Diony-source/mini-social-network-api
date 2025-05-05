package user

import (
	"errors"
	"mini-social-network-api/pkg/auth"
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
		return err
	}

	user := &User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: hashed,
	}

	return s.repo.CreateUser(user)
}

func (s *Service) Login(input LoginRequest) (*User, string, error) {
	user, err := s.repo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
