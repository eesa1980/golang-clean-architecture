package requesthandler

import (
	"clean-architecture/pkg/web/api/middlewares/request/docs-handler"
	. "clean-architecture/pkg/web/crosscutting"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func heartbeat() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello W")
	}
}

func New(app *fiber.App, deps *Dependencies) {
	app.Get("/heartbeat", heartbeat())
	app.Get("/swagger/docs", docshandler.New(deps))
	app.Get("/swagger/*", swagger.New(swagger.Config{URL: "/swagger/docs"}))
}
