package application

import (
	"fmt"
	"wire-demo-2/pkg/domain"
	"wire-demo-2/pkg/infrastructure"
)

type GetUserById func(id int) domain.User

func MakeGetUserById(deps infrastructure.Dependencies) GetUserById {
	return func(id int) domain.User {
		user, err := deps.UserService.GetUser(id)

		if err != nil {
			fmt.Print(err)
		}

		return user
	}
}
