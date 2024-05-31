package controller

import (
	"errors"
	"net/http"
	"oauth/enum"
	"oauth/model"
	"oauth/respMsg"
	"oauth/service"
	"oauth/sys"
	"oauth/utils"
	"strings"
)

func AccessToken(w http.ResponseWriter, r *http.Request) {
	var requestBody model.GetAccessTokenReq
	ParseRequestBody(w, r, &requestBody)

	sys.Logger().Debugf("Access Token request body: %+v", requestBody)

	if requestBody.GrantType == "" {
		HandleError(w, respMsg.INCORRECT_INPUT, "grant_type is required")
		return
	}

	// system secret contains systemCode
	systemSecret := r.Header.Get("system-secret")

	// TODO 如果未如果未來有做 client，需要額為驗證該 client 是否有准許請求的 grant_type
	var token model.AuthTokenResponse
	var err error
	switch requestBody.GrantType {
	case enum.PASSWORD.Val:
		if requestBody.AccountOrEmail == "" || requestBody.Password == "" {
			HandleError(w, respMsg.INCORRECT_INPUT, "account and password are required")
			return
		}
		token, err = service.GetAccessToken(requestBody.AccountOrEmail, requestBody.Password, systemSecret)
		break
	case enum.REFRESH_TOKEN.Val:
		tkn := r.Header.Get("Authorization")
		sys.Logger().Debugf("refresh access token: %s", tkn)
		token, err = service.RefreshAccessToken(tkn)
		break
	case enum.CLIENT_CREDENTIALS.Val:
		if requestBody.ClientID == "" || requestBody.ClientSecret == "" {
			HandleError(w, respMsg.INCORRECT_INPUT, "client_id and secret are required")
			return
		}
		token, err = service.ClientCredentials(requestBody.ClientID, requestBody.ClientSecret)
		break
	default:
		err = errors.New("grant_type not support")
		return
	}

	if err != nil {
		sys.Logger().Debugf("get access token failed, err:%v", err)
		HandleError(w, respMsg.WARNING, err.Error())
		return
	}

	HandleSuccess(w, token)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		HandleError(w, respMsg.UNAUTHORIZED, nil)
		return
	}

	accessToken := strings.TrimPrefix(auth, "Bearer ")
	authCode, _, _, err := utils.ValidateToken(accessToken)

	if err != nil {
		HandleError(w, respMsg.UNAUTHORIZED, nil)
		return
	}
	service.ExpireAccessToken(authCode)
	HandleSuccess(w, "success")
}
