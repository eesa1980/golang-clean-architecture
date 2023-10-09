package services

import (
	"github.com/google/wire"
	"wire-demo-2/pkg/application/common/interfaces"
	filehandler "wire-demo-2/pkg/infrastructure/service/file-handler"
)

type Services struct {
	FileHandler interfaces.IFileHandlerService
}

var ServiceSet = wire.NewSet(
	filehandler.New,
	wire.Struct(new(Services), "*"),
)
