package dao

import (
	"database/sql"
	"errors"
	"oauth/config"
	"oauth/entity"
)

func SelectEnableClientByClientID(clientID string) (*entity.Client, error) {
	db := config.GetDB()
	var client entity.Client
	query := "select id, client_id, client_name, client_secret, enable from oauth.client where client_id = ? and enable = 1"
	err := db.QueryRow(query, clientID).Scan(&client.ID, &client.ClientID, &client.ClientName, &client.ClientSecret, &client.Enable)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("client does not exist")
		}
		return nil, err
	}
	return &client, nil
}
