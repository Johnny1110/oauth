package controller

import (
	"errors"
	"net/http"
	"oauth/model"
	"oauth/respMsg"
	"oauth/service"
	"oauth/sys"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var requestBody model.CreateAccountReq
	ParseRequestBody(w, r, &requestBody)
	if requestBody.Email == "" && requestBody.Account == "" {
		HandleError(w, respMsg.INCORRECT_INPUT, errors.New("missing username or email"))
		return
	}
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

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var requestBody model.UpdateAccountPwdReq
	ParseRequestBody(w, r, &requestBody)
	if requestBody.AuthCode == "" || requestBody.NewPassword == "" {
		HandleError(w, respMsg.INCORRECT_INPUT, errors.New("missing authCode or newPassword"))
		return
	}
	err := service.UpdatePassword(requestBody.AuthCode, requestBody.NewPassword)
	if err != nil {
		sys.Logger().Warningf("fail to update password %v", err)
		HandleError(w, respMsg.WARNING, errors.New("failed to update password"))
		return
	}
	HandleSuccess(w, nil)
}
