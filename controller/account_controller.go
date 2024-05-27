package controller

import (
	"net/http"
	"oauth/model"
	"oauth/respMsg"
	"oauth/service"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody model.CreateAccountReq
	ParseRequestBody(w, r, &requestBody)

	authCode, err := service.CreateAccount(requestBody.Email, requestBody.Account, requestBody.Password)
	if err != nil {
		HandleError(w, respMsg.WARNING, err.Error())
		return
	}
	myMap := map[string]string{
		"authCode": authCode,
	}
	HandleSuccess(w, myMap)
}
