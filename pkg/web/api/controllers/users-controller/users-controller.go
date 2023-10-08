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
		user := deps.Application.Users.GetUserById(id)

		return ctx.Status(200).JSON(user)
	})

	app.Get("/users", func(ctx *fiber.Ctx) error {
		users := deps.Application.Users.ListUsers()

		return ctx.Status(200).JSON(users)
	})

}
