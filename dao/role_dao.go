package dao

import (
	"oauth/config"
	"oauth/entity"
	"oauth/sys"
)

func GetRoleDefaultScopes(roleName string) ([]entity.Scope, error) {
	db := config.GetDB()
	query := "SELECT scope.id, scope.resouce_id, scope.scope FROM oauth.scope as scope inner join oauth.role_default_scopes as rds on scope.id = rds.scope_id inner join oauth.role as role on rds.role_id = role.id where role.role_name = ?"
	rows, err := db.Query(query, roleName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scopes []entity.Scope
	for rows.Next() {
		var scope entity.Scope
		err := rows.Scan(&scope.ID, &scope.ResourceID, &scope.Scope)
		if err != nil {
			return nil, err
		}
		scopes = append(scopes, scope)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return scopes, nil
}

func SelectAccountRolesByAuthCode(authCode string) []entity.Role {
	db := config.GetDB()
	query := "SELECT role.id, role.role_name from oauth.role as role inner join oauth.account_roles as acc_roles on acc_roles.role_id = role.id inner join oauth.account as acc on acc.id = acc_roles.account_id where acc.auth_code = ?"
	var roles []entity.Role
	rows, err := db.Query(query, authCode)
	if err != nil {
		sys.Logger().Errorf("selectAccountRolesByAuthCode error: %s", err.Error())
		return roles
	}
	defer rows.Close()

	for rows.Next() {
		var role entity.Role
		err := rows.Scan(&role.ID, &role.RoleName)
		if err != nil {
			sys.Logger().Errorf("selectAccountRolesByAuthCode error: %s", err.Error())
			return roles
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		sys.Logger().Errorf("selectAccountRolesByAuthCode error: %s", err.Error())
		return roles
	}
	return roles
}
