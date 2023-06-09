package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var (
	RepositoryError = errors.New("unable to handle Repository request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func (r *repository) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users (id, email, password)
			VALUES ($1, $2,$3)`
	if user.Email == "" || user.Password == "" {
		return RepositoryError
	}
	_, err := r.db.ExecContext(ctx, sql, user.Id, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	sql := `SELECT email FROM users
			WHERE id = $1`
	err := r.db.QueryRow(sql, id).Scan(&email)
	if err != nil {
		return "", RepositoryError
	}
	return email, nil
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
