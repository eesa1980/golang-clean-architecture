package services

import (
	"github.com/google/wire"
	"wire-demo-2/pkg/application/common/interfaces"
	jsonfilehandler "wire-demo-2/pkg/infrastructure/service/json-file-handler"
)

type Services struct {
	CsvHandler interfaces.IJSONFileHandler
}

var ServiceSet = wire.NewSet(
	jsonfilehandler.New,
	wire.Struct(new(Services), "*"),
)
