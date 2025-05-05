package follow

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Follow(followerID, followeeID int64) error {
	query := `INSERT INTO follows (follower_id, followee_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, followerID, followeeID)
	return err
}
