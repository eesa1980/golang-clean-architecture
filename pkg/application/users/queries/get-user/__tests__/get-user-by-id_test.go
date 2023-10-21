package getuserbyidtest

import (
	getuserbyid "clean-architecture/pkg/application/users/queries/get-user"
	"testing"
)

func TestUserFound(t *testing.T) {

	getUser := getuserbyid.New(&MockUserRepository{})
	user := getUser(1)

	if user.ID != 1 {
		t.Errorf("Expected user id to be 1, got %v", user.ID)
	}
}

func TestUserNotFound(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	getUser := getuserbyid.New(&MockUserRepository{})
	getUser(2)
}
