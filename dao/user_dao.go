package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"oauth/config"
	"oauth/model"
)

// GetUserInfo fetches user information from the database by ID.
func GetUserInfo(id string) (model.User, error) {
	db := config.GetDB()
	var user model.User
	query := "SELECT id, account, password, enabled, locked ,expired ,auth_code,create_date, update_date  FROM basemall.account WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Account, &user.Password, &user.Enabled, &user.Locked, &user.Expired, &user.AuthCode, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// AddUserInfo adds a new user to the database.
func AddUserInfo(user model.User) error {
	db := config.GetDB()
	query := "INSERT INTO basemall.account (account, password, enabled, locked, expired, auth_code) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, user.Account, user.Password, user.Enabled, user.Locked, user.Expired, user.AuthCode)
	return err
}
