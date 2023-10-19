package services

import (
	"clean-architecture/pkg/application/common/interfaces"
	filehandler "clean-architecture/pkg/infrastructure/service/file-handler"
	"github.com/google/wire"
)

type Services struct {
	FileHandler interfaces.IFileHandlerService
}

var ServiceSet = wire.NewSet(
	filehandler.New,
	wire.Struct(new(Services), "*"),
)
