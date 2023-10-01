package application

import (
	. "wire-demo-2/pkg/domain"
	. "wire-demo-2/pkg/infrastructure"
)

type GetUserById func(id int) User

func MakeGetUserById(deps Dependencies) GetUserById {
	return func(id int) User {
		user, _ := deps.UserService.GetUser(id)
		return user
	}
}
