package follow

import (
	"errors"
	"mini-social-network-api/pkg/logger"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) FollowUser(followerID int64, input FollowRequest) error {
	logger.Log.WithFields(map[string]interface{}{
		"follower_id": followerID,
		"followee_id": input.FolloweeID,
	}).Info("attempting to follow user")

	if followerID == input.FolloweeID {
		logger.Log.Warn("user attempted to follow themselves")
		return errors.New("cannot follow yourself")
	}

	if err := s.repo.Follow(followerID, input.FolloweeID); err != nil {
		logger.Log.WithError(err).WithFields(map[string]interface{}{
			"follower_id": followerID,
			"followee_id": input.FolloweeID,
		}).Error("failed to follow user in repository")
		return err
	}

	logger.Log.WithFields(map[string]interface{}{
		"follower_id": followerID,
		"followee_id": input.FolloweeID,
	}).Info("user followed successfully")

	return nil
}
