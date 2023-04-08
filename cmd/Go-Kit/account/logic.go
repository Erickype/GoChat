package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func (s *service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	gId, _ := uuid.NewV4()
	id := gId.String()
	user := User{
		Id:       id,
		Email:    email,
		Password: password,
	}
	if err := s.repository.CreateUser(ctx, user); err != nil {
		err := level.Error(logger).Log("err", err)
		if err != nil {
			return "", err
		}
		return "", err
	}

	return "Success", nil
}

func (s *service) GetUser(ctx context.Context, id string) (string, error) {
	panic("implement me")
}

func NewService(repository Repository, logger log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}
