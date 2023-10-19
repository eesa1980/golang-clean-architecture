package listusers

import (
	"clean-architecture/pkg/domain"
	"clean-architecture/pkg/infrastructure"
)

type ListUsers func() []domain.User

func New(deps infrastructure.Dependencies) ListUsers {
	return func() []domain.User {
		users, _ := deps.UserRepository.ListUsers()
		return users
	}
}
