package infrastructure

import (
	"wire-demo-2/pkg/application/common/interfaces"
	"wire-demo-2/pkg/domain"
)

var mockUser = []domain.User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
	{ID: 3, Name: "Bob Smith"},
}

type userRepository struct {
	UserService interfaces.IUserService
}

func (u userRepository) GetUser(id int) (domain.User, error) {
	for _, user := range mockUser {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, nil
}

func (u userRepository) ListUsers() ([]domain.User, error) {
	return mockUser, nil
}

func MakeUserRepository() interfaces.IUserService {
	return userRepository{}
}
