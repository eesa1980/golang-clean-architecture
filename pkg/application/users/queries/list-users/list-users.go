package listusers

import (
	"wire-demo-2/pkg/domain"
	"wire-demo-2/pkg/infrastructure"
)

type ListUsers func() []domain.User

func New(deps infrastructure.Dependencies) ListUsers {
	return func() []domain.User {
		users, _ := deps.UserRepository.ListUsers()
		return users
	}
}
