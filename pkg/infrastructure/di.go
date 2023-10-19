package infrastructure

import (
	"clean-architecture/pkg/infrastructure/repository"
	"clean-architecture/pkg/infrastructure/service"
	"github.com/google/wire"
)

type Dependencies struct {
	repositories.Repositories
	services.Services
}

var Set = wire.NewSet(
	repositories.RepositorySet,
	services.ServiceSet,
	wire.Struct(new(Dependencies), "*"),
)
