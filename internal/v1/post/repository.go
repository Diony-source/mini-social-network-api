package post

import (
	"database/sql"
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
