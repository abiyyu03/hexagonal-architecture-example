package outbound

import (
	"go-projects/hexagonal-example/internal/adapter/outbound/repository"

	"go.uber.org/dig"
)

type Outbound struct {
	dig.In

	Repository repository.Repository
}

func Register(container *dig.Container) error {
	if err := repository.Register(container); err != nil {
		return err
	}

	return nil
}
