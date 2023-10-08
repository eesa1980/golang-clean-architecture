package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	. "wire-demo-2/pkg/web/api/controllers/users-controller"
	. "wire-demo-2/pkg/web/api/middlewares/request"
	. "wire-demo-2/pkg/web/api/middlewares/response"
	. "wire-demo-2/pkg/web/crosscutting"
)

type App struct {
	MakeApp func()
}

func MakeApp(
	deps *Dependencies,
) App {
	app := *fiber.New(
		fiber.Config{
			ErrorHandler: MakeHandleException(deps),
		})

	app.Use(recover.New())

	MakeRequest(&app, deps)
	MakeUsersController(&app, deps)
	MakeResponse(&app, deps)

	err := app.Listen(":3000")

	if err == nil {
		fmt.Println("Server is running on port 3000")
	} else {
		fmt.Print(err)
	}

	return App{}
}
