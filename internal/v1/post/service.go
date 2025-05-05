package post

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreatePost(input CreatePostRequest, userID int64) error {
	return s.repo.CreatePost(userID, input.Content)
}
