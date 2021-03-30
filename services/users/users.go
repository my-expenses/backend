package users

import (
	usersDBInteractions "backend/database/users"
	usersModel "backend/models/users"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (uint, error) {
	user, err := usersDBInteractions.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}
	credential := usersDBInteractions.GetCredentialsByUserID(user.ID)
	return user.ID, bcrypt.CompareHashAndPassword([]byte(credential.Password), []byte(password))
}

func NewUser(user *usersModel.User) error {
	return usersDBInteractions.NewUser(user)
}

func NewCredentials(credential *usersModel.Credential) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(credential.Password), 10)
	credential.Password = string(hashedPassword)
	return usersDBInteractions.NewCredentials(credential)
}