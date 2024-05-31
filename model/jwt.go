package model

import "github.com/golang-jwt/jwt/v4"

// Claims structure for JWT
type Claims struct {
	AuthCode   string   `json:"authCode"`
	ClientID   string   `json:"clientID"`
	ClientName string   `json:"clientName"`
	Email      string   `json:"email"`
	Roles      []string `json:"roles"`
	Scopes     []string `json:"scopes"`
	jwt.RegisteredClaims
}
