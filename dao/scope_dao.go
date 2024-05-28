package dao

import (
	"oauth/config"
	"oauth/entity"
	"oauth/sys"
)

func SelectAccountScopesByAuthCode(authCode string) []entity.Scope {
	db := config.GetDB()
	query := "SELECT scope.id, scope.resouce_id, scope.scope FROM oauth.scope as scope inner join oauth.account_scopes as acc_scopes on acc_scopes.scope_id = scope.id inner join oauth.account as acc on acc_scopes.account_id = acc.id where acc.auth_code = ? and acc_scopes.enable = true"
	var scopes []entity.Scope
	rows, err := db.Query(query, authCode)
	if err != nil {
		sys.Logger().Errorf("selectAccountRolesByAuthCode error: %s", err.Error())
		return scopes
	}
	defer rows.Close()

	for rows.Next() {
		var scope entity.Scope
		err := rows.Scan(&scope.ID, &scope.ResourceID, &scope.Scope)
		if err != nil {
			sys.Logger().Errorf("selectAccountRolesByAuthCode error: %s", err.Error())
			return scopes
		}
		scopes = append(scopes, scope)
	}

	if err = rows.Err(); err != nil {
		sys.Logger().Errorf("selectAccountRolesByAuthCode error: %s", err.Error())
		return scopes
	}
	return scopes
}
