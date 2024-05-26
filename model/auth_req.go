package model

type CreateAccountReq struct {
	Email    string `json:"email"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
