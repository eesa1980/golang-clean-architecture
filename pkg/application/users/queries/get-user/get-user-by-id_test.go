package getuserbyid

import (
	"clean-architecture/pkg/domain"
	mock_interfaces "clean-architecture/tests/mocks/interfaces"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func setup(t *testing.T) (GetUserById, *mock_interfaces.MockIUserRepository) {
	ctrl := gomock.NewController(t)
	userRepository := mock_interfaces.NewMockIUserRepository(ctrl)

	// Mock GetUser success
	userRepository.EXPECT().GetUser(1).DoAndReturn(func(id int) (domain.User, error) {
		time.Sleep(1 * time.Second)
		return domain.User{
			ID:        1,
			FirstName: "firstname",
			LastName:  "lastname",
			Email:     "user@example.com",
		}, nil
	}).AnyTimes()

	// Mock GetUser failure
	userRepository.EXPECT().GetUser(2).DoAndReturn(func(id int) (domain.User, error) {
		time.Sleep(1 * time.Second)
		return domain.User{}, errors.New(
			"error getting user")
	}).AnyTimes()

	getUserById := New(userRepository)

	return getUserById, userRepository
}

func TestUserFound(t *testing.T) {
	getUserById, _ := setup(t)

	user := getUserById(1)

	// Confirm that the user is returned
	if user.ID == 1 {
		t.Logf("User found: %v", user)
		return
	}

	t.Errorf("Expected user id to be 1, got %v", user.ID)

}

func TestUserNotFound(t *testing.T) {
	// Confirm that the function throws a panic when the user is not found
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to panic, but did not")
			return
		}

		t.Logf("User not found")
	}()

	getUserById, _ := setup(t)
	getUserById(2)

}
