package users_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	. "wire-demo-2/pkg/application/users/queries/get-user"
	. "wire-demo-2/pkg/application/users/queries/list-users"
	. "wire-demo-2/pkg/web/crosscutting"
)

type IUsersController interface {
	GetUserByIdController(getUserById GetUserById)
	ListUsersController(listUsers ListUsers)
}

type UsersController struct {
	usersController IUsersController
}

func (_ *UsersController) GetUserByIdController(getUserById GetUserById) {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetUserByIdController")
		idStr := r.URL.Query().Get("id")

		fmt.Println("GetUserByIdController")

		// log idStr to console
		fmt.Println(idStr)

		// parse sting to int
		id, _ := strconv.Atoi(idStr)

		user := getUserById(id)
		json.NewEncoder(w).Encode(user)
	})
}

func (_ *UsersController) ListUsersController(listUsers ListUsers) {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ListUsersController")
		users := listUsers()
		json.NewEncoder(w).Encode(users)
	})
}

func MakeUsersController(
	deps Dependencies,
) UsersController {
	usersController := &UsersController{}
	usersController.GetUserByIdController(deps.Application.Users.Queries.GetUserById)
	usersController.ListUsersController(deps.Application.Users.Queries.ListUsers)

	return UsersController{
		usersController,
	}
}
