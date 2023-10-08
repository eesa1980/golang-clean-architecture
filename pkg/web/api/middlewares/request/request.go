package requesthandler

import (
	"github.com/gofiber/fiber/v2"
	. "wire-demo-2/pkg/web/crosscutting"
)

func New(app *fiber.App, deps *Dependencies) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	app.Get("/api-docs", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Not Implemented yet")
	})
}
