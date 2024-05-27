package controller

import (
	"net/http"
	"oauth/model"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody model.CreateAccountReq
	ParseRequestBody(w, r, &requestBody)
	// TODO create user
	println("account: ", requestBody.Account)
	println("password: ", requestBody.Password)
	println("email: ", requestBody.Email)

	HandleSuccess(w, nil)
}
