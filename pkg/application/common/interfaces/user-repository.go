package interfaces

import (
	"clean-architecture/pkg/domain"
)

type IUserRepository interface {
	GetUser(id int) (domain.User, error)
	ListUsers() ([]domain.User, error)
}
