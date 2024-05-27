package middleware

import (
	"net/http"
	"oauth/controller"
	"oauth/respMsg"
	"oauth/route"
	"oauth/sys"
	"oauth/utils"
	"strings"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sys.Logger().Debugf("Started %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		sys.Logger().Debugf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpMethod := r.Method
		url := r.URL.Path

		for _, route := range route.Routes {
			if route.Method == httpMethod && route.Pattern == url {
				if route.NeedAuth {
					auth := r.Header.Get("Authorization")
					sys.Logger().Debugf("auth checking %s %s auth-contect: [%s]", httpMethod, url, auth)
					if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
						sys.Logger().Warningf("auth failed: %s %s", httpMethod, url)
						controller.HandleError(w, respMsg.UNAUTHORIZED, nil)
						return
					}

					authToken := strings.TrimPrefix(auth, "Bearer ")
					// Implement token validation and role/scope checking here
					userRoles, userScopes, err := utils.ValidateToken(authToken)

					if err != nil {
						sys.Logger().Warningf("auth failed: %s %s - invalid token", httpMethod, url)
						controller.HandleError(w, respMsg.UNAUTHORIZED, nil)
						return
					}
					if !utils.HasRequiredPermissions(userRoles, route.Permission.Roles, userScopes, route.Permission.Scopes) {
						sys.Logger().Warningf("auth failed: %s %s - insufficient permissions", httpMethod, url)
						controller.HandleError(w, respMsg.UNAUTHORIZED, nil)
						return
					}
				} else {
					sys.Logger().Debugf("api not set oauth protection: %s %s", httpMethod, url)
				}

			}
		}

		next.ServeHTTP(w, r)
	})
}
