package middleware

import (
	"net/http"
	"oauth/controller"
	"oauth/respMsg"
	"oauth/route"
	"oauth/sys"
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

					if auth != "Bearer valid-token" {
						sys.Logger().Warningf("auth failed: %s %s", httpMethod, url)
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
