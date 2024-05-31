package dao

import (
	"oauth/config"
	"oauth/entity"
	"oauth/sys"
)

func SelectAccountScopesByAuthCode(authCode string) []entity.Scope {
	db := config.GetDB()
	query := "SELECT scope.id, scope.resouce_id, scope.scope FROM oauth.scope as scope inner join oauth.account_scopes as acc_scopes on acc_scopes.scope_id = scope.id inner join oauth.account as acc on acc_scopes.account_id = acc.id where acc.auth_code = ? and acc_scopes.enable = 1"
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

func SelectClientScopesByClientPK(clientPK int) []entity.Scope {
	db := config.GetDB()
	query := "SELECT scope.id, scope.resouce_id, scope.scope FROM oauth.scope as scope inner join oauth.client_scopes as cs on scope.id = cs.scope_id where cs.client_id = ? and cs.enable = 1 and scope.enable = 1;"
	var scopes []entity.Scope
	rows, err := db.Query(query, clientPK)
	if err != nil {
		sys.Logger().Errorf("SelectClientScopesByClientPK error: %s", err.Error())
		return scopes
	}
	defer rows.Close()

	for rows.Next() {
		var scope entity.Scope
		err := rows.Scan(&scope.ID, &scope.ResourceID, &scope.Scope)
		if err != nil {
			sys.Logger().Errorf("SelectClientScopesByClientPK error: %s", err.Error())
			return scopes
		}
		scopes = append(scopes, scope)
	}

	if err = rows.Err(); err != nil {
		sys.Logger().Errorf("SelectClientScopesByClientPK error: %s", err.Error())
		return scopes
	}
	return scopes
}
