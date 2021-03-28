package users

import (
	usersDBInteractions "backend/database/users"
	usersModel "backend/models/users"
	customErrors "backend/services/errors"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(user *usersModel.User, password, confirmPassword string) error {
	if password != confirmPassword {
		return &customErrors.PasswordsDontMatchError{}
	}
	return usersDBInteractions.NewUser(user)
}

func NewCredentials(credential *usersModel.Credential) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(credential.Password), 10)
	credential.Password = string(hashedPassword)
	return usersDBInteractions.NewCredentials(credential)
}