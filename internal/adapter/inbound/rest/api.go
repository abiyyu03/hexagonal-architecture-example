package rest

import "github.com/gofiber/fiber/v2"

func (i Inbound) ApiRoutes(app *fiber.App) {
	app.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK LURD")
	})

	app.Post("/user/register", i.User.RegisterUser)
}
