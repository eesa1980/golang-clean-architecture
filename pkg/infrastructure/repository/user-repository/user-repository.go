package userrepository

import (
	"clean-architecture/pkg/application/common/interfaces"
	"clean-architecture/pkg/domain"
	"encoding/json"
	"errors"
)

type userRepository struct {
	interfaces.IFileHandlerService
}

// getUsersFromJson is a helper function to get users from a json file
func getUsersFromJson(u *userRepository) *[]domain.User {
	file, _ := u.LoadFile("mock-user-data.json")
	decoder := json.NewDecoder(file)

	var users *[]domain.User
	err := decoder.Decode(&users)

	if err != nil {
		panic(err)
	}

	defer u.Close()

	return users
}

// GetUser is a method on the userRepository struct that returns a user by id
func (u *userRepository) GetUser(id int) (domain.User, error) {
	var users []domain.User = *getUsersFromJson(u)

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return domain.User{}, errors.New("User not found")

}

// ListUsers is a method on the userRepository struct that returns a list of users
func (u *userRepository) ListUsers() ([]domain.User, error) {
	var users []domain.User = *getUsersFromJson(u)
	return users, nil
}

func New(fileHandlerService interfaces.IFileHandlerService) interfaces.IUserRepository {
	return &userRepository{
		fileHandlerService,
	}
}
