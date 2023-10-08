package web

import (
	"github.com/gofiber/fiber/v2"
	. "wire-demo-2/pkg/application/common/exceptions"
)

func MakeHandleNotFound() func(ctx *fiber.Ctx) error {

	return func(ctx *fiber.Ctx) error {
		return ctx.Status(404).JSON(Exception{
			Message:     "Unable to find the requested resource",
			StatusCode:  404,
			Title:       "Not Found",
			Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		})
	}
}
