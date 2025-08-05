package rest

import (
	"go-projects/hexagonal-example/internal/adapter/inbound/rest/user"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	User user.Handler
}
