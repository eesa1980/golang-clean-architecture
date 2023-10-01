package application

import (
	"github.com/google/wire"
	. "wire-demo-2/pkg/application/users"
)

type Dependencies struct {
	Users
}

var Set = wire.NewSet(
	UsersSet,
	wire.Struct(new(Dependencies), "*"),
)
