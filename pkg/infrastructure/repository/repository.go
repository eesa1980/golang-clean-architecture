package repositories

import (
	"clean-architecture/pkg/application/common/interfaces"
	"clean-architecture/pkg/application/common/structs"
	"clean-architecture/pkg/infrastructure/config"
	"clean-architecture/pkg/infrastructure/repository/user-repository"
	"github.com/google/wire"
)

type Repositories struct {
	Config         structs.ApplicationConfig
	UserRepository interfaces.IUserRepository
}

var RepositorySet = wire.NewSet(
	confighandler.New,
	userrepository.New,
	wire.Struct(new(Repositories), "*"),
)
