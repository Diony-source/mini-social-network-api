package user

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(user *User) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, user.Username, user.Email, user.PasswordHash)
	return err
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	var user User
	query := `SELECT id, username, email, password_hash, created_at FROM users WHERE email=$1`
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
