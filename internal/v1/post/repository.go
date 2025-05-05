package post

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreatePost(authorID int64, content string) error {
	query := `INSERT INTO posts (author_id, content) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, authorID, content)
	return err
}
