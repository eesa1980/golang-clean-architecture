package application

import (
	"wire-demo-2/pkg/domain"
)

type IUserService interface {
	GetUser(id int) (domain.User, error)
	ListUsers() ([]domain.User, error)
}
