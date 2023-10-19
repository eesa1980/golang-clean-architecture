package container

import (
	"clean-architecture/pkg/application"
	"clean-architecture/pkg/infrastructure"
)

type Dependencies struct {
	Application    application.Dependencies
	Infrastructure infrastructure.Dependencies
}
