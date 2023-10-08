package responsehandler

import (
	"github.com/gofiber/fiber/v2"
	"wire-demo-2/pkg/web/api/middlewares/response/not-found-handler"
	"wire-demo-2/pkg/web/crosscutting"
)

func New(app *fiber.App, deps *container.Dependencies) {
	app.Get("*", notfoundhandler.New())
}
