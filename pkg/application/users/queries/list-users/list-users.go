package application

import (
	. "wire-demo-2/pkg/domain"
	. "wire-demo-2/pkg/infrastructure"
)

type ListUsers func() []User

func MakeListUsers(deps Dependencies) ListUsers {
	return func() []User {
		users, _ := deps.UserService.ListUsers()
		return users
	}
}
