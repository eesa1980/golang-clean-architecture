package repositories

import (
	"clean-architecture/pkg/application/common/interfaces"
	"clean-architecture/pkg/infrastructure/repository/user-repository"
	"github.com/google/wire"
)

type Repositories struct {
	UserRepository interfaces.IUserRepository
}

var RepositorySet = wire.NewSet(
	userrepository.New,
	wire.Struct(new(Repositories), "*"),
)
