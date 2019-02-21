package controller

import (
	"github.com/caarlos0/apps-operator/pkg/controller/app"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, app.Add)
}
