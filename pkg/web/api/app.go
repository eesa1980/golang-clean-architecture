package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"wire-demo-2/pkg/web/api/controllers/users-controller"
	requesthandler "wire-demo-2/pkg/web/api/middlewares/request"
	"wire-demo-2/pkg/web/api/middlewares/response"
	"wire-demo-2/pkg/web/api/middlewares/response/exception-handler"
	"wire-demo-2/pkg/web/crosscutting"
)

type App struct {
	MakeApp func()
}

func New(
	deps *container.Dependencies,
) App {
	app := *fiber.New(
		fiber.Config{
			ErrorHandler: exceptionhandler.New(),
		})

	app.Use(recover.New())

	requesthandler.New(&app, deps)
	userscontroller.New(&app, deps)
	responsehandler.New(&app, deps)

	err := app.Listen(":3000")

	if err == nil {
		fmt.Println("Server is running on port 3000")
	} else {
		fmt.Print(err)
	}

	return App{}
}
