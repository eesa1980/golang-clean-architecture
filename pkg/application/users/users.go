package application

import (
	"github.com/google/wire"
	. "wire-demo-2/pkg/application/users/queries"
)

type Users struct {
	Queries
}

var UsersSet = wire.NewSet(
	QueriesSet,
	wire.Struct(new(Users), "*"),
)
