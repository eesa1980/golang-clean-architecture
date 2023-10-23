package listusers

import (
	"clean-architecture/pkg/domain"
	mock_interfaces "clean-architecture/tests/mocks/interfaces"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func setup(t *testing.T) (ListUsers, *mock_interfaces.MockIUserRepository) {
	ctrl := gomock.NewController(t)
	userRepository := mock_interfaces.NewMockIUserRepository(ctrl)

	listUsers := New(userRepository)

	return listUsers, userRepository
}

func TestUsersFound(t *testing.T) {
	listUsers, userRepository := setup(t)

	userRepository.EXPECT().ListUsers().DoAndReturn(func() ([]domain.User, error) {
		time.Sleep(1 * time.Second)
		return []domain.User{
			{
				ID:        1,
				FirstName: "firstname1",
				LastName:  "lastname1",
				Email:     "user1@example.com",
			}, {
				ID:        2,
				FirstName: "firstname2",
				LastName:  "lastname2",
				Email:     "user2@example.com",
			}}, nil
	}).AnyTimes()

	users := listUsers()

	usersLength := len(users)

	if usersLength == 2 {
		t.Logf("Users found: %v", users)
		return
	}

	t.Errorf("Expected users length to be 2, got %v", usersLength)

}

func TestUserNotFound(t *testing.T) {
	listUsers, userRepository := setup(t)

	userRepository.EXPECT().ListUsers().DoAndReturn(func() ([]domain.User, error) {
		time.Sleep(1 * time.Second)
		return []domain.User{}, errors.New("error getting users")
	}).AnyTimes()

	users := listUsers()
	usersLength := len(users)

	if usersLength == 0 {
		t.Logf("Users not found")
		return
	}

}
