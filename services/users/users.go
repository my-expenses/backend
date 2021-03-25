package users

import (
	usersDBInteractions "backend/database/users"
	usersModel "backend/models/users"
)

func NewUser(user *usersModel.User, password string) error {
	return usersDBInteractions.NewUser(user, &usersModel.Credential{
		Password: password,
	})
}