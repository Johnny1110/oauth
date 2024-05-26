package main

import (
	"fmt"
	"net/http"
	"oauth/config"
	"oauth/middleware"
	"oauth/route"
	"oauth/sys"
)

func DispatchHandler(w http.ResponseWriter, r *http.Request) {
	route.Dispatch(w, r)
}

func main() {
	mux := http.NewServeMux()
	dispatchHandler := http.HandlerFunc(DispatchHandler)

	// wrapper by log and auth handler
	loggedHandler := middleware.LoggingMiddleware(dispatchHandler)
	authenticatedHandler := middleware.AuthenticationMiddleware(loggedHandler)

	// handle all request
	mux.Handle("/", authenticatedHandler)

	properties := config.GetProperties()
	port := properties.Port
	addr := fmt.Sprintf(":%s", port)
	sys.Logger().Debug("oauth server listening on port: ", port)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		sys.Logger().Error("Could not start server: ", err)
		return
	}
}
