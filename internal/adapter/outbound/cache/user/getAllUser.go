package user

import (
	"context"
	"encoding/json"
	"go-projects/hexagonal-example/internal/adapter/outbound/model"
)

func (c userCache) GetAllUser(ctx context.Context) ([]model.User, error) {
	var key = "users:getAll"
	result := c.Package.Cache.Client.Get(ctx, key)

	var users []model.User
	str, err := result.Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(str), &users); err != nil {
		return nil, err
	}

	return users, nil
}

type IGetAllUser interface {
	GetAllUser(ctx context.Context) ([]model.User, error)
}
