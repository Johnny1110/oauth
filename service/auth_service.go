package service

import (
	"errors"
	"oauth/model"
	"oauth/sys"
	"oauth/utils"
)

func GetAccessToken(accountOrEmail, password, systemSecret string) (model.AuthTokenResponse, error) {
	// TODO 1. parse systemSecret to get systemType
	systemType, err := parseSystemSecret(systemSecret)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}
	sys.Logger().Debugf("[GetAccessToken] accountOrEmail: %s, systemType: %s", accountOrEmail, systemType)

	var hashedPassword string
	//TODO 2. get User from db by accountOrEmail and systemType (btw check locked, expired, enable).
	if !(accountOrEmail == "root" || accountOrEmail == "Jarvan1110@gmail.com") {
		return model.AuthTokenResponse{}, errors.New("account or email is invalid")
	}

	// compare password.
	if !utils.CheckPassword(password, hashedPassword) {
		return model.AuthTokenResponse{}, errors.New("password is invalid.")
	}

	// TODO 3. if already have cached accessToken, then refreshAccessToken() and return new Token

	// TODO for the test : dummy user roles and scopes data
	authCode := "USNDQ23110NS"
	username := "Johnny"
	email := "Jarvan1110@gmail.com"
	userRoles := []string{"ROLE_USER_L1", "ROLE_USER_L2"}
	userScopes := []string{"cart.write", "cart.read", "mall.read", "consumer.read", "consumer.write", "order.write", "order.read"}

	// 4. generate jwt token with roles and scopes.
	accessToken, err := utils.GenerateJWT(authCode, email, username, userRoles, userScopes, 3600)
	refreshToken, err := utils.GenerateJWT(authCode, email, username, userRoles, userScopes, 3600)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}

	response := model.AuthTokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresIn:  3600,
		RefreshTokenExpiresIn: 3600,
		TokenType:             "Bearer",
		AuthCode:              authCode,
	}

	// 5.TODO cached accessToken and refreshToken
	return response, nil
}

// TODO
func parseSystemSecret(secret string) (string, error) {
	if secret == "" {
		return "", errors.New("system secret is invalid.")
	}
	return "FRIZO_USER", nil
}
