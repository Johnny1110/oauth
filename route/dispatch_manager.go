package route

import "net/http"

func Dispatch(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method
	url := r.URL.Path

	for _, route := range Routes {
		if route.Method == httpMethod && route.Pattern == url {
			route.HandlerFunc(w, r)
			return
		}
	}

	http.NotFound(w, r)
}
