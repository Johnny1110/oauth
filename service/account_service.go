package service

import (
	"errors"
	"oauth/dao"
	"oauth/entity"
	"oauth/enum"
	"oauth/sys"
	"oauth/utils"
	"strings"
)

func CreateAccount(email, username, password string) (string, error) {
	systemType := enum.FRIZO_USER // default common user account
	hashedPassword := utils.HashPassword(password)

	account := dao.SelectAccountByEmailOrUsername(systemType.Code, email, username)

	if account.AuthCode != "" {
		sys.Logger().Warningf("account %s already exists", email)
		return "", errors.New("account already exist")
	}

	authCode := utils.GenRandomString(24, false)
	defaultRole := enum.ROLE_SYS_USER_L1
	roleDefaultScopes, err := dao.GetRoleDefaultScopes(defaultRole.Val)
	if err != nil {
		sys.Logger().Warningf("get role default scopes failed: %s", err.Error())
		return "", err
	}

	account = entity.Account{
		SystemID:     systemType.Code,
		AuthCode:     authCode,
		Email:        strings.ToLower(email),
		Username:     strings.ToLower(username),
		PasswordHash: hashedPassword,
		Enable:       true,
		Locked:       false,
		Expired:      false,
	}

	accountID, err := dao.InsertAccount(account)
	if err != nil {
		sys.Logger().Warningf("insert account failed: %s", err.Error())
		return "", err
	}

	rowEffectRow, err := dao.InsertAccountRole(accountID, defaultRole.Code)
	if err != nil || rowEffectRow == 0 {
		sys.Logger().Warningf("insert role failed: %s", err.Error())
		return "", err
	}
	scopeEffectRow, err := dao.InsertAccountScopes(accountID, roleDefaultScopes)
	if err != nil || scopeEffectRow == 0 {
		sys.Logger().Warningf("insert scopes failed: %s", err.Error())
		return "", err
	}
	return authCode, err
}

func UpdatePassword(authCode string, newPassword string) error {
	systemType := enum.FRIZO_USER // default common user account
	hashedPassword := utils.HashPassword(newPassword)
	account, err := dao.SelectAccountByAuthCode(systemType.Code, authCode)
	if err != nil {
		return err
	}
	if account.AuthCode == "" {
		sys.Logger().Warningf("account %s not exists", authCode)
		return errors.New("account not exist")
	}
	account.PasswordHash = hashedPassword
	effectRow, err := dao.UpdateAccount(account)
	if err != nil {
		return err
	}
	if effectRow == 0 {
		return errors.New("update account failed")
	}
	return nil
}
