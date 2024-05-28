package model

type CreateAccountReq struct {
	Email    string `json:"email"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UpdateAccountPwdReq struct {
	AuthCode    string `json:"authCode"`
	NewPassword string `json:"newPassword"`
}

type LockAccountReq struct {
	Email   string `json:"email"`
	Account string `json:"account"`
}

type UnLockAccountReq struct {
	Email   string `json:"email"`
	Account string `json:"account"`
}
