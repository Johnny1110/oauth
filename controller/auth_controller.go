package controller

import (
	"net/http"
	"oauth/model"
	"oauth/respMsg"
	"oauth/service"
	"oauth/sys"
)

func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	var requestBody model.GetAccessTokenReq
	ParseRequestBody(w, r, &requestBody)
	// system secret contains systemCode
	systemSecret := r.Header.Get("system-secret")

	token, err := service.GetAccessToken(requestBody.AccountOrEmail, requestBody.Password, systemSecret)

	if err != nil {
		sys.Logger().Debugf("get access token failed, err:%v", err)
		HandleError(w, respMsg.WARNING, "account or password incorrect.")
		return
	}

	HandleSuccess(w, token)
}
