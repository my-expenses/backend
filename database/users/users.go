package users

import (
	dbInstance "backend/database"
	customErrors "backend/database/errors"
	usersModel "backend/models/users"
	"strings"
)

func GetUserByEmail(email string) (*usersModel.User, error) {
	var user usersModel.User
	dbInstance.GetDBConnection().Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return nil, &customErrors.NoUserFoundError{}
	}
	return &user, nil
}

func GetCredentialsByUserID(userID uint) *usersModel.Credential {
	var credential usersModel.Credential
	dbInstance.GetDBConnection().Where("user_id = ?", userID).Find(&credential)
	return &credential
}

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
