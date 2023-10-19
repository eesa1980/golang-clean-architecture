package responsehandler

import (
	"clean-architecture/pkg/web/api/middlewares/response/not-found-handler"
	"clean-architecture/pkg/web/crosscutting"
	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App, deps *container.Dependencies) {
	app.Get("*", notfoundhandler.New())
}
