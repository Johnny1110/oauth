package model

type GetAccessTokenReq struct {
	AccountOrEmail string `json:"accountOrEmail"`
	Password       string `json:"password"`
}
