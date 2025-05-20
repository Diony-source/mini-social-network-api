package post

import (
	"database/sql"
	"errors"
	"mini-social-network-api/pkg/logger"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatePost(userID int64, content string) error {
	query := `INSERT INTO posts (author_id, content) VALUES ($1, $2)`

	if _, err := r.db.Exec(query, userID, content); err != nil {
		logger.Log.WithError(err).WithFields(map[string]interface{}{
			"user_id": userID,
			"query":   "insert post",
		}).Error("failed to execute post insert query")
		return err
	}

	return nil
}

func (r *Repository) GetByID(postID int64) (*Post, error) {
	query := `SELECT id, author_id, content, created_at FROM posts WHERE id = $1`
	post := &Post{}

	if err := r.db.QueryRow(query, postID).Scan(&post.ID, &post.AuthorID, &post.Content, &post.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		logger.Log.WithError(err).WithField("post_id", postID).Error("failed to execute post select query")
		return nil, err
	}
	return post, nil
}

func (r *Repository) Update(postID int64, content string) error {
	query := `UPDATE posts SET content = $1 WHERE id = $2`
	_, err := r.db.Exec(query, content, postID)
	if err != nil {
		logger.Log.WithError(err).WithField("post_id", postID).Error("failed to execute post update query")
	}
	return err
}
