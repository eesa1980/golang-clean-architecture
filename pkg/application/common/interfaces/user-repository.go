package interfaces

import (
	"wire-demo-2/pkg/domain"
)

type IUserRepository interface {
	GetUser(id int) (domain.User, error)
	ListUsers() ([]domain.User, error)
}
