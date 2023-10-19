package getuserbyid

import (
	. "clean-architecture/pkg/application/common/exception"
	"clean-architecture/pkg/domain"
	"clean-architecture/pkg/infrastructure"
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
