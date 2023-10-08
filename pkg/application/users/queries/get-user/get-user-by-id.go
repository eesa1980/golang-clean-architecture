package application

import (
	. "wire-demo-2/pkg/application/common/exceptions"
	"wire-demo-2/pkg/domain"
	"wire-demo-2/pkg/infrastructure"
)

type GetUserById func(id int) domain.User

func MakeGetUserById(deps infrastructure.Dependencies) GetUserById {

	return func(id int) domain.User {

		user, err := deps.UserService.GetUser(id)

		if err != nil {
			panic(&NotFoundException{
				Message: err.Error(),
			})
		}

		return user
	}
}
