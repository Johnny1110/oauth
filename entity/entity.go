package entity

type Scope struct {
	ID         int
	ResourceID int    `json:"resource_id"`
	Scope      string `json:"scope"`
	Enable     bool   `json:"enable"`
	Desc       string `json:"desc"`
}

type Account struct {
	ID           int    `json:"-"`
	SystemID     int    `json:"system_id"`
	AuthCode     string `json:"auth_code"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Enable       bool   `json:"enable"`
	Locked       bool   `json:"locked"`
	Expired      bool   `json:"expired"`
}
