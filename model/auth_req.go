package model

type GetAccessTokenReq struct {
	AccountOrEmail string `json:"account"`
	Password       string `json:"password"`
	GrantType      string `json:"grant_type"`
}
