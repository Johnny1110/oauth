package main

import (
	"fmt"
	"net/http"
	"oauth/config"
	"oauth/controller"
)

func main() {
	properties := config.GetProperties()
	port := properties.Port
	addr := fmt.Sprintf(":%s", port)
	http.HandleFunc("/oauth/user", controller.UserInfo)
	fmt.Printf("oauth server listening at port: %v\n", port)
	http.ListenAndServe(addr, nil)
}
