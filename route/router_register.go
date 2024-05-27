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
	Permission  AuthPermission
}

type AuthPermission struct {
	Roles  []string
	Scopes []string
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
		Permission: AuthPermission{
			Roles:  []string{"ROLE_ADMIN"},
			Scopes: []string{"oauth.super", "oauth.write"},
		},
	},

	{
		Method:      "POST",
		Pattern:     "/oauth/token",
		HandlerFunc: controller.GetAccessToken,
		NeedAuth:    false,
	},
}
