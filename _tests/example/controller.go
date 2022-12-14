package example

import (
	"github.com/wolesgo/core/_tests/example/app/controller/authentication"
	"github.com/wolesgo/core/_tests/example/app/controller/home"
	"github.com/wolesgo/core/controller"
)

func Controllers(controllers *controller.ControllerCollections) {
	controllers.Add(home.NewHomeController)
	controllers.Add(authentication.NewLoginController)
	controllers.Add(authentication.NewLogoutController)
}
