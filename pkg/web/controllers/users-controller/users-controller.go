package web

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	. "wire-demo-2/pkg/web/crosscutting"
)

func MakeUsersController(
	app *fiber.App,
	deps *Dependencies,
) {
	app.Get("/users/:id", func(ctx *fiber.Ctx) error {
		id, _ := strconv.Atoi(ctx.Params("id"))
		user := deps.Application.GetUserById(id)
		return ctx.JSON(user)
	})

	app.Get("/users", func(ctx *fiber.Ctx) error {
		users := deps.Application.ListUsers()
		return ctx.JSON(users)
	})

}
