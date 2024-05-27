package dao

import (
	"oauth/config"
	"oauth/entity"
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
