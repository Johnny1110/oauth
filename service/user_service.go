package service

import (
	"oauth/dao"
	"oauth/model"
)

// GetUserInfo retrieves user information by ID.
func GetUserInfo(id string) (model.User, error) {
	return dao.GetUserInfo(id)
}

// AddUserInfo adds a new user to the database.
func AddUserInfo(user model.User) error {
	return dao.AddUserInfo(user)
}
