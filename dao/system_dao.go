package dao

import (
	"database/sql"
	"fmt"
	"oauth/config"
	"oauth/entity"
)

func GetSystemBySecret(secretString string) (*entity.System, error) {
	db := config.GetDB()
	var system entity.System
	query := "SELECT id, system_code, secret from oauth.system where secret = ?;"
	err := db.QueryRow(query, secretString).Scan(&system.ID, &system.SystemCode, &system.Secret)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no system found with the given secret")
		}
		return nil, err
	}
	return &system, nil
}
