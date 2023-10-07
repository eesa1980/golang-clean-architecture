package infrastructure

import (
	"errors"
	. "wire-demo-2/pkg/application/common/interfaces"
	. "wire-demo-2/pkg/domain"
)

var mockUsers = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
	{ID: 3, Name: "Bob Smith"},
}

type userRepository struct {
	UserService IUserService
}

func (u userRepository) GetUser(id int) (User, error) {
	for _, user := range mockUsers {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")

}

func (u userRepository) ListUsers() ([]User, error) {
	return mockUsers, nil
}

func MakeUserRepository() IUserService {
	return userRepository{}
}
