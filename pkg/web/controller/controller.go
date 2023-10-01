package web

import (
	. "wire-demo-2/pkg/web/controller/users-controller"
	. "wire-demo-2/pkg/web/crosscutting"
)

type Controller struct {
	usersController UsersController
}

func MakeController(
	deps Dependencies,
) Controller {
	usersController := MakeUsersController(deps)

	return Controller{
		usersController,
	}
}
