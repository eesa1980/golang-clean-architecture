package app

import (
	userscontroller "clean-architecture/pkg/web/api/controllers/users-controller"
	requesthandler "clean-architecture/pkg/web/api/middlewares/request"
	responsehandler "clean-architecture/pkg/web/api/middlewares/response"
	exceptionhandler "clean-architecture/pkg/web/api/middlewares/response/exception-handler"
	container "clean-architecture/pkg/web/crosscutting"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(
	deps *container.Dependencies,
) *fiber.App {

	// Initialize Fiber app with custom error handler middleware
	app := *fiber.New(
		fiber.Config{
			AppName:               deps.Infrastructure.ApplicationConfig.AppName,
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
