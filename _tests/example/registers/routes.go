package registers

import (
	router "github.com/wolesgo/core/router/console"
)

func ConsoleRouter(router *router.Router) {
	router.Command(
		"home",
		"To login into the system",
	).Controller(
		"_tests/example/app/controllers/home/HomeController",
		"Action",
	)

	router.Command(
		"login",
		"To login into the system",
	).Controller(
		"_tests/example/app/controllers/LoginController",
		"Action",
	)

	router.Command(
		"logout",
		"To logout from the system",
	).Controller(
		"_tests/example/app/controllers/LoginController",
		"Action",
	)
}
