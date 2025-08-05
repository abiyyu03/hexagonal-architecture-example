package user

import (
	"context"
	"go-projects/hexagonal-example/internal/adapter/inbound/rest/user/port"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) RegisterUser(fctx *fiber.Ctx) error {
	var ctx = context.Background()
	var (
		response port.BaseResponse
		request  port.CreateUserRequest
	)

	if err := fctx.BodyParser(&request); err != nil {
		return err
	}

	err := h.Service.User.CreateUser(ctx, request.ToServiceUser())
	if err != nil {
		return err
	}

	return fctx.Status(fiber.StatusCreated).JSON(
		response.ToResponse(
			"User created successfully",
			fiber.StatusCreated,
			nil,
			nil,
		),
	)
}
