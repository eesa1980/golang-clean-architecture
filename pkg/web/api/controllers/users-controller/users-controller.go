package userscontroller

import (
	_ "clean-architecture/pkg/application/common/exception"
	_ "clean-architecture/pkg/domain"
	. "clean-architecture/pkg/web/crosscutting"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// getUserByIdController
//
//	@Router			/api/v1/users/{id} [get]
//	@Summary		Get user by ID
//	@Description	Get user by ID
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	domain.User
//	@Failure		400	{object}	exception.Exception
//	@Failure		404	{object}	exception.Exception
//	@Tags			Users
func getUserByIdController(
	deps *Dependencies,
) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		id, _ := strconv.Atoi(ctx.Params("id"))
		user := deps.Application.Users.GetUserById(id)

		return ctx.Status(200).JSON(user)
	}
}

// listUsersController
//
//	@Router			/api/v1/users [get]
//	@Summary		Get all users
//	@Description	Get all users
//	@Success		200	{object}	[]domain.User
//	@Failure		400	{object}	exception.Exception
//	@Failure		500	{object}	exception.Exception
//	@Tags			Users
func listUsersController(
	deps *Dependencies,
) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		users := deps.Application.Users.ListUsers()

		return ctx.Status(200).JSON(users)
	}
}

func New(
	app *fiber.App,
	deps *Dependencies,
) {
	v1 := app.Group("/api/v1/users")

	v1.Get("/:id", getUserByIdController(deps))
	v1.Get("/test/123", listUsersController(deps))
}
