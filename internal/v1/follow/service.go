package follow

import "errors"

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) FollowUser(followerID int64, input FollowRequest) error {
	if followerID == input.FolloweeID {
		return errors.New("cannot follow yourself")
	}

	return s.repo.Follow(followerID, input.FolloweeID)
}
