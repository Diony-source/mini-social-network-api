package post

import "mini-social-network-api/pkg/logger"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreatePost(input CreatePostRequest, userID int64) error {
	logger.Log.WithFields(map[string]interface{}{
		"user_id": userID,
		"content": input.Content,
	}).Info("attempting to create post")

	if err := s.repo.CreatePost(userID, input.Content); err != nil {
		logger.Log.WithError(err).Error("failed to create post in repository")
		return err
	}

	logger.Log.WithField("user_id", userID).Info("post created in service layer")
	return nil
}
