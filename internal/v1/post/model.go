package post

type Post struct {
	ID        int64  `json:"id"`
	AuthorID  int64  `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type CreatePostRequest struct {
	Content string `json:"content" validate:"omitempty"`
}

type UpdatePostRequest struct {
	Content string `json:"content" validate:"omitempty"`
}
