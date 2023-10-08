package application

import (
	"github.com/google/wire"
	"wire-demo-2/pkg/application/users/queries/get-user"
	"wire-demo-2/pkg/application/users/queries/list-users"
)

type Queries struct {
	getuserbyid.GetUserById
	listusers.ListUsers
}

var QueriesSet = wire.NewSet(
	getuserbyid.New,
	listusers.New,
	wire.Struct(new(Queries), "*"),
)
