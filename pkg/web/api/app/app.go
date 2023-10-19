package app

import (
	"clean-architecture/pkg/web/api/controllers/users-controller"
	"clean-architecture/pkg/web/api/middlewares/request"
	"clean-architecture/pkg/web/api/middlewares/response"
	"clean-architecture/pkg/web/api/middlewares/response/exception-handler"
	"clean-architecture/pkg/web/crosscutting"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(
	deps *container.Dependencies,
) *fiber.App {

	// Initialize Fiber app with custom error handler middleware
	app := *fiber.New(
		fiber.Config{
			AppName:               deps.Infrastructure.Config.AppName,
			ErrorHandler:          exceptionhandler.New(),
			DisableStartupMessage: false,
		})

	// Prevents panic exceptions from stopping the application
	app.Use(recover.New())

	// Register middlewares and routes
	requesthandler.New(&app, deps)
	userscontroller.New(&app, deps)
	responsehandler.New(&app, deps)

	return &app
}
