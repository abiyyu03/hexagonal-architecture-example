package user

import (
	"context"
	"go-projects/hexagonal-example/internal/adapter/outbound/model"
)

func (r user) CreateUser(ctx context.Context, user model.User) error {
	if err := r.Package.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return nil
}

type ICreateUser interface {
	CreateUser(ctx context.Context, user model.User) error
}
