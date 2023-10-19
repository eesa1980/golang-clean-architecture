package application

import (
	. "clean-architecture/pkg/application/users/queries"
	"github.com/google/wire"
)

type Users struct {
	Queries
}

var UsersSet = wire.NewSet(
	QueriesSet,
	wire.Struct(new(Users), "*"),
)
