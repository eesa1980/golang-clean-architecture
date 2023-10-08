package infrastructure

import (
	"github.com/google/wire"
	"wire-demo-2/pkg/infrastructure/repository"
	"wire-demo-2/pkg/infrastructure/service"
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
