package dao

import (
	"oauth/config"
	"oauth/entity"
)

func InsertAccount(account entity.Account) (int64, error) {
	sql := "INSERT INTO oauth.account(system_id, auth_code, email, username, password_hash) VALUES (?, ?, ?, ?, ?)"
	result, err := config.GetDB().Exec(sql, account.SystemID, account.AuthCode, account.Email, account.Username, account.PasswordHash)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func InsertAccountRole(accountID int64, roleID int) (int64, error) {
	sql := "INSERT INTO oauth.account_roles (account_id, role_id, enable) VALUES (?, ?, true)"
	result, err := config.GetDB().Exec(sql, accountID, roleID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func InsertAccountScopes(accountID int64, scopes []entity.Scope) (int64, error) {
	if len(scopes) == 0 {
		return 0, nil
	}

	// 构建批量插入的 SQL 语句
	sql := "INSERT INTO oauth.account_scopes (account_id, scope_id, enable) VALUES "
	values := []interface{}{}

	for _, scope := range scopes {
		sql += "(?, ?, true),"
		values = append(values, accountID, scope.ID)
	}

	sql = sql[:len(sql)-1]

	result, err := config.GetDB().Exec(sql, values...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func SelectAccountByEmailOrUsername(systemTypeId int, email string, username string) entity.Account {
	db := config.GetDB()
	var account entity.Account
	query := "SELECT id,system_id,auth_code,email,username,password_hash from oauth.account where system_id = ? and (email = ? or username = ?)"
	db.QueryRow(query, systemTypeId, email, username).Scan(&account.ID, &account.SystemID, &account.AuthCode, &account.Email, &account.Username, &account.PasswordHash)
	return account
}
