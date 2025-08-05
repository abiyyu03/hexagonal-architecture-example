package user

import (
	"context"
	"go-projects/hexagonal-example/internal/service/model"
)

func (s service) CreateUser(ctx context.Context, user model.User) error {
	err := s.Repository.User.CreateUser(ctx, user.ToOutbondUser())
	if err != nil {
		return err
	}

	return nil
}

type ICreateUser interface {
	CreateUser(ctx context.Context, user model.User) error
}
