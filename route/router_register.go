package route

import (
	"net/http"
	"oauth/controller"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	NeedAuth    bool
}

var Routes = []Route{
	{
		Method:      "GET",
		Pattern:     "/oauth/healthcheck",
		HandlerFunc: controller.HealthCheck,
		NeedAuth:    false,
	},

	{
		Method:      "POST",
		Pattern:     "/oauth/account",
		HandlerFunc: controller.CreateAccount,
		NeedAuth:    true,
	},
}
