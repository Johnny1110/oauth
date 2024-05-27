package model

type CreateAccountReq struct {
	Email    string `json:"email"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UpdateAccountPwdReq struct {
	Email       string `json:"email"`
	Account     string `json:"account"`
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
