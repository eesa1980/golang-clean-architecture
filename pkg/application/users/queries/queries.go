package application

import (
	"github.com/google/wire"
	. "wire-demo-2/pkg/application/users/queries/get-user"
	. "wire-demo-2/pkg/application/users/queries/list-users"
)

type Queries struct {
	GetUserById GetUserById
	ListUsers   ListUsers
}

var QueriesSet = wire.NewSet(
	MakeGetUserById,
	MakeListUsers,
	wire.Struct(new(Queries), "*"),
)
