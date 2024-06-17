package infrastructure

import (
	"clean-architecture/pkg/application/common/structs"
	confighandler "clean-architecture/pkg/infrastructure/config"
	repositories "clean-architecture/pkg/infrastructure/repository"
	services "clean-architecture/pkg/infrastructure/service"

	"github.com/google/wire"
)

type Dependencies struct {
	repositories.Repositories
	services.Services
	structs.ApplicationConfig
}

var Set = wire.NewSet(
	confighandler.New,
	repositories.RepositorySet,
	services.ServiceSet,
	wire.Struct(new(Dependencies), "*"),
)
