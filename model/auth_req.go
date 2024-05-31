package model

type GetAccessTokenReq struct {
	AccountOrEmail string `json:"account"`
	Password       string `json:"password"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"secret"`
	GrantType      string `json:"grant_type"`
}
