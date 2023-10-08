package repositories

import (
	"github.com/google/wire"
	"wire-demo-2/pkg/application/common/interfaces"
	"wire-demo-2/pkg/infrastructure/repository/user-repository"
)

type Repositories struct {
	UserRepository interfaces.IUserRepository
}

var RepositorySet = wire.NewSet(
	userrepository.New,
	wire.Struct(new(Repositories), "*"),
)
