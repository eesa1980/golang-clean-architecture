package getuserbyidtest

import (
	"clean-architecture/pkg/domain"
	"errors"
)

type MockUserRepository struct{}

func (u *MockUserRepository) ListUsers() ([]domain.User, error) {
	panic("implement me")
}

func (u *MockUserRepository) GetUser(id int) (domain.User, error) {
	if id != 1 {
		return domain.User{}, errors.New("user not found")
	}

	return domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "random@email.com",
	}, nil
}
