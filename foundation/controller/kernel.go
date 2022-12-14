package foundation

import (
	"github.com/wolesgo/core/container"
	"github.com/wolesgo/core/contracts"
	"github.com/wolesgo/core/controller"
)

type Controller struct {
	app container.Container

	controllers *controller.ControllerCollections
}

func NewController(app container.Container, opt controller.ControllerOption) contracts.ControllerKernelContract {
	if opt.BaseModulePath != "" {
		opt.BaseModulePath = app.GetBaseModulePath(opt.BaseModulePath)
	}

	return Controller{
		app:         app,
		controllers: controller.NewCollections(opt),
	}
}

func (controller Controller) GetApplication() container.Container {
	return controller.app
}

func (controller Controller) GetControllers() *controller.ControllerCollections {
	return controller.controllers
}

func (controller Controller) bootstrap() {
}

func (controller Controller) Register() {
}
