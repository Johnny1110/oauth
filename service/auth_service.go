package service

import (
	"errors"
	"oauth/cache"
	"oauth/dao"
	"oauth/entity"
	"oauth/enum"
	"oauth/model"
	"oauth/sys"
	"oauth/utils"
	"strings"
)

func GetAccessToken(accountOrEmail, password, systemSecret string) (model.AuthTokenResponse, error) {
	//  1. parse systemSecret to get systemCode
	systemID, err := parseSystemSecretToID(systemSecret)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}
	sys.Logger().Debugf("[GetAccessToken] accountOrEmail: %s, systemID: %s", accountOrEmail, systemID)

	account, _ := dao.SelectAccountByHybridParams(systemID, accountOrEmail)
	if account == nil {
		return model.AuthTokenResponse{}, errors.New("account or email is invalid")
	}

	if !account.Enable && account.Locked && account.Expired {
		return model.AuthTokenResponse{}, errors.New("account or email is locked")
	}

	// compare password.
	if !utils.CheckPassword(password, account.PasswordHash) {
		return model.AuthTokenResponse{}, errors.New("password is invalid.")
	}

	return generateAuthResponse(account)
}

func generateAuthResponse(account *entity.Account) (model.AuthTokenResponse, error) {
	sys.Logger().Debugf("[generateAuthResponse] account: %v", account)
	authCode := account.AuthCode
	username := account.Username
	email := account.Email
	userRoles := dao.SelectAccountRolesByAuthCode(authCode)
	userRoleNames := make([]string, len(userRoles))
	for i, role := range userRoles {
		userRoleNames[i] = role.RoleName
	}
	userScopes := dao.SelectAccountScopesByAuthCode(authCode)
	userScopeNames := make([]string, len(userScopes))
	for i, scope := range userScopes {
		userScopeNames[i] = scope.Scope
	}
	// for refresh token
	refreshTokenScopeNames := []string{"oauth.refresh"}

	// token expireTime timeUnit = min
	accessTokenExpiresIn := 3600
	refreshTokenExpiresIn := 3600
	// 4. generate jwt token with roles and scopes.
	accessToken, err := utils.GenerateJWT(authCode, email, username, userRoleNames, userScopeNames, accessTokenExpiresIn)
	refreshToken, err := utils.GenerateJWT(authCode, email, username, userRoleNames, refreshTokenScopeNames, refreshTokenExpiresIn)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}

	response := model.AuthTokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresIn:  accessTokenExpiresIn,
		RefreshTokenExpiresIn: refreshTokenExpiresIn,
		TokenType:             "Bearer",
		AuthCode:              authCode,
	}

	// 5 cached accessToken and refreshToken
	cache.Set(cache.GenKeyString(enum.ACCESS_TOKEN_RDS_KEY, authCode), accessToken, accessTokenExpiresIn)
	cache.Set(cache.GenKeyString(enum.REFRESH_TOKEN_RDS_KEY, authCode), refreshToken, refreshTokenExpiresIn)
	return response, nil
}

// get systemCode by secret
func parseSystemSecretToID(secretString string) (int, error) {
	if secretString == "" {
		return 0, errors.New("system secret string is invalid.")
	}
	system, _ := dao.GetSystemBySecret(secretString)
	if system == nil {
		return 0, errors.New("system secret is invalid.")
	} else {
		return system.ID, nil
	}
}

func RefreshAccessToken(bearerToken string, systemSecret string) (model.AuthTokenResponse, error) {
	//  1. parse systemSecret to get systemCode
	systemID, err := parseSystemSecretToID(systemSecret)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}

	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
		return model.AuthTokenResponse{}, errors.New("refresh token is invalid.")
	}

	refreshToken := strings.TrimPrefix(bearerToken, "Bearer ")
	// Implement token validation and role/scope checking here
	authCode, _, userScopes, err := utils.ValidateToken(refreshToken)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}

	checkScopeFlag := false
	for _, scope := range userScopes {
		if scope == "oauth.refresh" {
			checkScopeFlag = true
		}
	}

	if !checkScopeFlag {
		return model.AuthTokenResponse{}, errors.New("refresh token is invalid.")
	}
	sys.Logger().Debugf("select acc by systemID:%v authCode: %s", systemID, authCode)
	account, _ := dao.SelectAccountByAuthCode(systemID, authCode)
	if account == nil {
		return model.AuthTokenResponse{}, errors.New("authCode or systemSecret is invalid.")
	}
	return generateAuthResponse(account)
}
