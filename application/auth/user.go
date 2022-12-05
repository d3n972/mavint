package auth

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	IAuthenticate
	password []byte
	username string
}

func (u User) GetPassword() ([]byte, error) {
	if u.password != nil {
		return u.password, nil
	}
	return nil, ErrorUserNotInitialized
}
func (u User) SetPassword(passwd []byte) (bool, error) {
	password, err := bcrypt.GenerateFromPassword(passwd, BCryptCost)
	if err != nil {
		return false, err
	}
	u.password = password
	return true, nil
}
func (u User) GetUsername() (string, error) {
	if u.username != "" {
		return u.username, nil
	}
	return "", ErrorUserNotInitialized
}
func (u User) SetUsername(username string) (bool, error) {
	u.username = username
	return true, nil
}
func (u User) Authenticate(pass []byte) (bool, error) {
	pwHash, _ := u.GetPassword()
	if bcrypt.CompareHashAndPassword(pwHash, pass) != nil {
		return true, nil
	}
	return false, ErrorWrongPassword
}
