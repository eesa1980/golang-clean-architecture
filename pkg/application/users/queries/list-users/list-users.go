package listusers

import (
	"clean-architecture/pkg/application/common/interfaces"
	"clean-architecture/pkg/domain"
)

type ListUsers func() []domain.User
type IUserRepository interfaces.IUserRepository

func New(userRepository interfaces.IUserRepository) ListUsers {
	return func() []domain.User {
		users, _ := userRepository.ListUsers()

		return users
	}
}
