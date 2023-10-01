package crosscutting

import (
	"wire-demo-2/pkg/application"
	"wire-demo-2/pkg/infrastructure"
)

type Dependencies struct {
	Application    application.Dependencies
	Infrastructure infrastructure.Dependencies
}
