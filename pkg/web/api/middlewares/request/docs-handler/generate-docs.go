package docshandler

import (
	. "clean-architecture/pkg/web/crosscutting"
	"github.com/gofiber/fiber/v2"
)

func New(deps *Dependencies) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		file, err := deps.Infrastructure.Services.FileHandler.LoadFile("pkg/web/api/api-docs/swagger.json")

		if err != nil {
			panic(err)
		}

		json, err := deps.Infrastructure.Services.FileHandler.ToJson()

		if err != nil {
			panic(err)
		}

		err = file.Close()

		if err != nil {
			panic(err)
		}

		return ctx.Status(200).JSON(&json)

	}
}
