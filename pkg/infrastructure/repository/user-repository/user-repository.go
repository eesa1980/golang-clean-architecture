package userrepository

import (
	"errors"
	"wire-demo-2/pkg/application/common/interfaces"
	"wire-demo-2/pkg/domain"
)

type dependencies struct {
	interfaces.IJSONFileHandler
}

func readUsers(d *dependencies) *[]domain.User {
	jsonUsers, _ := d.Load()
	var users []domain.User
	err := jsonUsers.Decode(&users)

	if err != nil {
		panic(err)
	}

	defer d.Close()

	return &users
}

func (deps *dependencies) GetUser(id int) (domain.User, error) {
	var users []domain.User = *readUsers(deps)

	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	defer deps.Close()

	return domain.User{}, errors.New("user not found")

}

func (deps *dependencies) ListUsers() ([]domain.User, error) {
	var users []domain.User = *readUsers(deps)
	return users, nil
}

func New(jsonFileHandler interfaces.IJSONFileHandler) interfaces.IUserRepository {
	return &dependencies{
		jsonFileHandler,
	}
}
