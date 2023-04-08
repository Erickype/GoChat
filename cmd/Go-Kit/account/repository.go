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
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetUser(ctx context.Context, id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
