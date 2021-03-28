package users

import (
	dbInstance "backend/database"
	customErrors "backend/database/errors"
	usersModel "backend/models/users"
	"strings"
)

func NewUser(user *usersModel.User) error {
	err := dbInstance.GetDBConnection().Create(user).Error
	if err != nil {
		strings.HasPrefix(err.Error(), "Error 1062")
		return &customErrors.DuplicateEmailError{}
	}
	return err
}

func NewCredentials(credentials *usersModel.Credential) error {
	return dbInstance.GetDBConnection().Create(credentials).Error
}
