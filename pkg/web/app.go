package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	. "wire-demo-2/pkg/web/controllers/users-controller"
	. "wire-demo-2/pkg/web/crosscutting"
	. "wire-demo-2/pkg/web/middlewares/request"
)

type App struct {
	MakeApp func()
}

func MakeApp(
	deps *Dependencies,
) App {
	app := *fiber.New()
	OnRequest(&app, deps)
	MakeUsersController(&app, deps)
	err := app.Listen(":3000")

	if err == nil {
		fmt.Println("Server is running on port 3000")
	} else {
		fmt.Print(err)
	}

	return App{}
}
