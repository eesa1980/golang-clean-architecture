package web

import (
	"github.com/gofiber/fiber/v2"
	. "wire-demo-2/pkg/web/crosscutting"
)

func OnResponse(app *fiber.App, deps *Dependencies) {
	app.Get("*", MakeHandleNotFound())
}
