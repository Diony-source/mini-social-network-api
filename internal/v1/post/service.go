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

func (s *Service) UpdatePost(input UpdatePostRequest, postID int64) error {
	logger.Log.WithFields(map[string]interface{}{
		"post_id": postID,
		"content": input.Content,
	}).Info("attempting to update post")

	if err := s.repo.Update(postID, input.Content); err != nil {
		logger.Log.WithError(err).Error("failed to update post in repository")
		return err
	}

	logger.Log.WithField("post_id", postID).Info("post updated in service layer")
	return nil
}

func (s *Service) GetByID(postID int64) (*Post, error) {
	logger.Log.WithField("post_id", postID).Info("attempting to get post by ID")

	post, err := s.repo.GetByID(postID)
	if err != nil {
		logger.Log.WithError(err).Error("failed to get post from repository")
		return nil, err
	}

	logger.Log.WithField("post_id", postID).Info("post retrieved successfully")
	return post, nil
}

func (s *Service) DeletePost(postID int64) error {
	logger.Log.WithField("post_id", postID).Info("attempting to delete post")

	if err := s.repo.Delete(postID); err != nil {
		logger.Log.WithError(err).Error("failed to delete post from repository")
		return err
	}

	logger.Log.WithField("post_id", postID).Info("post deleted successfully")
	return nil
}
