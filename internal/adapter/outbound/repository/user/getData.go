package user

import (
	"context"
	"go-projects/hexagonal-example/internal/adapter/outbound/model"
)

func (r user) GetUser(ctx context.Context) (model.User, error) {
	var response model.User
	if err := r.Package.DB.WithContext(ctx).Find(&response).Error; err != nil {
		return response, err
	}

	return response, nil
}

type IGetUser interface {
	GetUser(ctx context.Context) (model.User, error)
}
