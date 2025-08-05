package user

import (
	"context"
	"go-projects/hexagonal-example/internal/adapter/outbound/model"
)

func (r user) GetAll(ctx context.Context) ([]model.User, error) {
	var user []model.User
	if err := r.Package.DB.WithContext(ctx).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

type IGetAll interface {
	GetAll(ctx context.Context) ([]model.User, error)
}
