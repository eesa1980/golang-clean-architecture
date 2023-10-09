package getuserbyid

import (
	. "wire-demo-2/pkg/application/common/exception"
	"wire-demo-2/pkg/domain"
	"wire-demo-2/pkg/infrastructure"
)

type GetUserById func(id int) domain.User

func New(deps infrastructure.Dependencies) GetUserById {

	return func(id int) domain.User {

		user, err := deps.UserRepository.GetUser(id)

		if err != nil {
			panic(NewException[NotFound](err.Error()))
		}

		return user
	}
}
