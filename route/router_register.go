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
		Pattern:     "/oauth/account", // create account
		HandlerFunc: controller.CreateAccount,
		NeedAuth:    true,
		Permission: AuthPermission{
			Roles:  []string{"ROLE_SYS_ADMIN"},
			Scopes: []string{"oauth.sp", "oauth.write"},
		},
	},

	{
		Method:      "POST",
		Pattern:     "/oauth/token", // get access_token
		HandlerFunc: controller.AccessToken,
		NeedAuth:    false,
	},

	{
		Method:      "PUT",
		Pattern:     "/oauth/password", // update password 不開放一般 user 自己來改密碼（因為驗證手續 ex: email 驗證 需要在 consumer 模組做）
		HandlerFunc: controller.UpdatePassword,
		NeedAuth:    true,
		Permission: AuthPermission{
			Roles:  []string{"ROLE_SYS_ADMIN"},
			Scopes: []string{"oauth.sp", "oauth.write"},
		},
	},
}
