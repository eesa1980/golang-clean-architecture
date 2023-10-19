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

type App struct {
}

//	@title			Clean Architecture API
//	@version		1.0
//	@description	This is an example server for a Clean Architecture API written in Go.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @externalDocs	description="Find out more about Clean Architecture" url="www.google.com"

func New(
	deps *container.Dependencies,
) App {

	// Initialize Fiber app with custom error handler middleware
	app := *fiber.New(
		fiber.Config{
			ErrorHandler: exceptionhandler.New(),
		})

	// Prevents panic exceptions from stopping the application
	app.Use(recover.New())

	// Register middlewares
	requesthandler.New(&app, deps)
	userscontroller.New(&app, deps)
	responsehandler.New(&app, deps)

	err := app.Listen("localhost:3000")

	if err != nil {
		defer app.Shutdown()
		// stop the application
		panic(err)
	}

	return App{}
}
