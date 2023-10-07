//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-demo-2/pkg/application"
	"wire-demo-2/pkg/infrastructure"
	. "wire-demo-2/pkg/web/api"
	. "wire-demo-2/pkg/web/crosscutting"
)

var apps = wire.NewSet(
	wire.Struct(new(Dependencies), "*"),
	application.Set,
	infrastructure.Set,
	MakeApp,
)

func Initialize() App {
	wire.Build(apps)
	return App{}
}
