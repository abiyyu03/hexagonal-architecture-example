package service

import (
	"go-projects/hexagonal-example/internal/service/user"

	"go.uber.org/dig"
)

type Service struct {
	dig.In

	User user.UserService
}

func Register(container *dig.Container) error {
	if err := container.Provide(user.New); err != nil {
		return err
	}

	return nil
}
