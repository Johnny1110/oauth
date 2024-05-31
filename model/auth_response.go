package model

type AuthTokenResponse struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	AccessTokenExpiresIn  int    `json:"access_token_expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	AuthCode              string `json:"auth_code"`
	ClientID              string `json:"client_id"`
}
