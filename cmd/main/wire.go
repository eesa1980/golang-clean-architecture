//go:build wireinject
// +build wireinject

package main

import (
	"clean-architecture/pkg/application"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/web/api/server"
	"clean-architecture/pkg/web/crosscutting"
	"github.com/google/wire"
)

var apps = wire.NewSet(
	wire.Struct(new(container.Dependencies), "*"),
	application.Set,
	infrastructure.Set,
	server.New,
)

func Initialize() server.Server {
	wire.Build(apps)
	return server.Server{}
}
