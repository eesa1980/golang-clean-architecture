package server

import (
	"clean-architecture/pkg/web/api/app"
	container "clean-architecture/pkg/web/crosscutting"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/log"
)

type Server struct {
}

//	@title			Clean Architecture API
//	@version		1.0
//	@description	This is an example server for a Clean Architecture API written in Go.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
// @externalDocs	description="Find out more about Clean Architecture" url="www.google.com"
func New(
	deps *container.Dependencies,
) Server {
	newApp := app.New(deps)

	config := deps.Infrastructure.ApplicationConfig

	url := fmt.Sprintf(":%s", config.Port)
	log.Infof("Application is now running on http://127.0.0.1%s/swagger", url)

	err := newApp.Listen(url)

	if err != nil {
		defer func(newApp *fiber.App) {
			err := newApp.Shutdown()
			if err != nil {
				panic(err)
			}
		}(newApp)
		// stop the application
		panic(err)
	}

	return Server{}
}
