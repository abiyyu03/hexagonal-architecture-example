package user

import (
	"context"
	"go-projects/hexagonal-example/internal/service/model"
)

func (s service) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := s.Repository.User.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []model.User
	for _, value := range users {
		response = append(response, model.User{
			ID:    value.ID,
			Email: value.Email,
			Name:  value.Name,
		})
	}

	return response, nil
}

type IGetAll interface {
	GetAll(ctx context.Context) ([]model.User, error)
}
