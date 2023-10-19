package application

import (
	. "clean-architecture/pkg/application/users"
	"github.com/google/wire"
)

type Dependencies struct {
	Users
}

var Set = wire.NewSet(
	UsersSet,
	wire.Struct(new(Dependencies), "*"),
)
