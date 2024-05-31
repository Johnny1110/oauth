package service

import (
	"encoding/json"
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
		return model.AuthTokenResponse{}, errors.New("password is invalid")
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
	ExpireAccessToken(authCode)
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
	accountJson, err := json.Marshal(account)
	if err != nil {
		return model.AuthTokenResponse{}, errors.New("client json marshal failed")
	}
	cache.Set(cache.GenKeyString(enum.ACCESS_TOKEN_RDS_KEY, authCode+":"+accessToken), string(accountJson), accessTokenExpiresIn)
	cache.Set(cache.GenKeyString(enum.REFRESH_TOKEN_RDS_KEY, authCode+":"+refreshToken), string(accountJson), refreshTokenExpiresIn)
	return response, nil
}

// get systemCode by secret
func parseSystemSecretToID(secretString string) (int, error) {
	if secretString == "" {
		return 0, errors.New("system secret string is invalid")
	}
	system, _ := dao.GetSystemBySecret(secretString)
	if system == nil {
		return 0, errors.New("system secret is invalid")
	} else {
		return system.ID, nil
	}
}

func RefreshAccessToken(bearerToken string) (model.AuthTokenResponse, error) {
	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
		return model.AuthTokenResponse{}, errors.New("refresh token is invalid")
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
		return model.AuthTokenResponse{}, errors.New("refresh token is invalid")
	}
	ExpireAccessToken(authCode)
	sys.Logger().Debugf("select acc by authCode: %s", authCode)
	account, _ := dao.SelectAccountByAuthCode(authCode)
	if account == nil {
		return model.AuthTokenResponse{}, errors.New("authCode is invalid")
	}
	return generateAuthResponse(account)
}

func ClientCredentials(clientID string, clientSecret string) (model.AuthTokenResponse, error) {
	client, err := dao.SelectEnableClientByClientID(clientID)
	if client == nil || err != nil {
		return model.AuthTokenResponse{}, err
	}
	if client.ClientSecret != clientSecret {
		return model.AuthTokenResponse{}, errors.New("client secret is invalid")
	}

	clientRoles := dao.SelectClientRolesByClientPK(client.ID)
	roleNames := make([]string, len(clientRoles))
	for i, role := range clientRoles {
		roleNames[i] = role.RoleName
	}
	clientScopes := dao.SelectClientScopesByClientPK(client.ID)
	scopeNames := make([]string, len(clientScopes))
	for i, scope := range clientScopes {
		scopeNames[i] = scope.Scope
	}
	// token expireTime timeUnit = min
	accessTokenExpiresIn := 3600
	// 4. generate jwt token with roles and scopes.
	accessToken, err := utils.GenerateClientJWT(clientID, client.ClientName, roleNames, scopeNames, accessTokenExpiresIn)
	if err != nil {
		return model.AuthTokenResponse{}, err
	}

	response := model.AuthTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresIn: accessTokenExpiresIn,
		TokenType:            "Bearer",
		ClientID:             clientID,
	}

	// 5 cached accessToken and refreshToken
	clientJson, err := json.Marshal(client)
	if err != nil {
		return model.AuthTokenResponse{}, errors.New("client json marshal failed")
	}
	ExpireAccessToken(clientID)
	cache.Set(cache.GenKeyString(enum.ACCESS_TOKEN_RDS_KEY, clientID+":"+accessToken), string(clientJson), accessTokenExpiresIn)
	return response, nil
}

func ExpireAccessToken(authCode string) {
	accessTokenStoreKey := cache.GenKeyString(enum.ACCESS_TOKEN_RDS_KEY, authCode)
	refreshTokenStoreKey := cache.GenKeyString(enum.REFRESH_TOKEN_RDS_KEY, authCode)
	cache.ScanDelete(accessTokenStoreKey)
	cache.ScanDelete(refreshTokenStoreKey)
	sys.Logger().Debugf("[ExpireAccessToken] done.")
}
