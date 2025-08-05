package repository

import (
	"go-projects/hexagonal-example/internal/adapter/outbound/repository/user"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	User user.Repository
}

func Register(container *dig.Container) error {
	if err := container.Provide(user.New); err != nil {
		return err
	}

	return nil
}
