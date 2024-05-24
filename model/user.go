package model

// User represents a user in the system.
type User struct {
	ID        int    `json:"id"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Enabled   bool   `json:"enabled"`
	Locked    bool   `json:"locked"`
	Expired   bool   `json:"expired"`
	AuthCode  string `json:"auth_code"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
