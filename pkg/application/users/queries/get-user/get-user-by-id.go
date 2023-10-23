package getuserbyid

import (
	. "clean-architecture/pkg/application/common/exception"
	"clean-architecture/pkg/application/common/interfaces"
	"clean-architecture/pkg/domain"
)

type GetUserById func(id int) domain.User
type IUserRepository interfaces.IUserRepository

func New(userRepository interfaces.IUserRepository) GetUserById {

	return func(id int) domain.User {

		user, err := userRepository.GetUser(id)

		if err != nil {
			panic(NewException[NotFound](err.Error()))
		}

		return user
	}
}
