package application

import (
	"clean-architecture/pkg/application/users/queries/get-user"
	"clean-architecture/pkg/application/users/queries/list-users"
	"github.com/google/wire"
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
