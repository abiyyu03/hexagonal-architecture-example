package user

import (
	"context"
	"go-projects/hexagonal-example/internal/service/entity"
)

func (s service) CreateUser(ctx context.Context, user entity.User) error {
	err := s.Repository.User.CreateUser(ctx, user.ToOutbondUser())
	if err != nil {
		return err
	}

	return nil
}

type ICreateUser interface {
	CreateUser(ctx context.Context, user entity.User) error
}
