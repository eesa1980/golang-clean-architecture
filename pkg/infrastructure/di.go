package infrastructure

import (
	"github.com/google/wire"
	. "wire-demo-2/pkg/application/common/interfaces"
	. "wire-demo-2/pkg/infrastructure/repository"
)

type Dependencies struct {
	UserService IUserService
}

var Set = wire.NewSet(
	MakeUserRepository,
	wire.Struct(new(Dependencies), "*"),
)
