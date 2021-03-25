package users

import (
	dbInstance "backend/database"
	customErrors "backend/database/errors"
	usersModel "backend/models/users"
	"strings"
)

func NewUser(user *usersModel.User, credential *usersModel.Credential) error {
	err := dbInstance.GetDBConnection().Create(user).Error
	if err != nil {
		strings.HasPrefix(err.Error(), "Error 1062")
		return &customErrors.DuplicateEmailError{}
	}
	credential.UserID = user.ID
	return dbInstance.GetDBConnection().Create(credential).Error
}
