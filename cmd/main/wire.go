//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"net/http"
	"wire-demo-2/pkg/application"
	"wire-demo-2/pkg/infrastructure"
	controller "wire-demo-2/pkg/web/controller"
	. "wire-demo-2/pkg/web/crosscutting"
)

type App struct {
	MakeApp func()
}

func MakeApp(
	deps Dependencies,
) App {
	controller.MakeController(deps)
	http.ListenAndServe("localhost:3000", nil)
	return App{}
}

var apps = wire.NewSet(wire.Struct(new(Dependencies), "*"), controller.MakeController, application.Set, infrastructure.Set, MakeApp)

func Initialize() App {
	wire.Build(apps)
	return App{}
}
